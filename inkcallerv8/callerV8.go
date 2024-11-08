// Package inkcallerv8 is an implementation of inkcaller using ink js and V8 to run the JS.
package inkcallerv8

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
	"rogchap.com/v8go"
)

type (
	impl struct {
		prepareInkContext func(ctx context.Context, engineFilePath string, storyFilePath string) (*v8go.Context, error)
		stories           map[string]fileCache
		storiesMutex      *sync.RWMutex
		engines           map[string]fileCache
		enginesMutex      *sync.RWMutex
	}

	//20220207 performance: compiling is not worth it.
	//same perf, more allocs, code more complicated
	//I keep the code, hoping in the long run it is worth it (longer story file, better cache hit ratio in prod)

	//20220207 performance: same for caching file content in map.
	//benchmark is showing no diff compare to reading it each time.

	fileCache struct {
		content   string
		codeCache *v8go.CompilerCachedData
	}

	CallRes struct {
		Choices []Choice
		Text    string
		State   string
	}
	Choice struct {
		Text  string `json:"text"`
		Index uint16 `json:"index"`
	}
)

const inkVirtualJSFilePath = "s.js"

var _ inkcallerlib.InkCaller = (*impl)(nil)

func NewInkCallerV8() *impl {
	impl := &impl{
		stories:      make(map[string]fileCache),
		storiesMutex: &sync.RWMutex{},
		engines:      make(map[string]fileCache),
		enginesMutex: &sync.RWMutex{},
	}

	// creates a new JavaScript VM
	iso1 := v8go.NewIsolate()

	impl.prepareInkContext = func(ctx context.Context, engineFilePath string, storyFilePath string) (*v8go.Context, error) {
		// new context within the VM
		inkCtx := v8go.NewContext(iso1)

		//performance: loading+compiling the 2 files in parallel is worse.
		//(using an errgroup for the benchmark)
		//60% of speed + more alloc

		err := impl.loadFileInContext(iso1, inkCtx, engineFilePath, impl.enginesMutex, impl.engines, nil)
		if err != nil {
			return nil, err
		}
		err = impl.loadFileInContext(iso1, inkCtx, storyFilePath, impl.storiesMutex, impl.stories,
			func(path, sFile string) string {
				//this is the difference between the .json and .js
				//this is small enough we can be compatible with both.
				if strings.HasSuffix(storyFilePath, ".json") {
					sFile = fmt.Sprintf("var storyContent=%s", sFile)
				}
				return sFile
			})
		if err != nil {
			return nil, err
		}
		return inkCtx, nil
	}

	return impl
}

func (c *impl) Call(
	ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	opts ...inkcallerlib.InkCallerOptionsFunc,
) (*inkcallerlib.InkCallerOutput, error) {
	inkCtx, err := c.prepareInkContext(ctx, engineFilePath, storyFilePath)
	if err != nil {
		return nil, err
	}
	options := inkcallerlib.ReadOptions(opts)

	execJs := c.prepareJS(options)
	val, err := inkCtx.RunScript(execJs, inkVirtualJSFilePath)
	if err != nil {
		var JSError *v8go.JSError
		if errors.As(err, &JSError) {
			//Special formatting with Message+Location+StackTrace when using %+v
			return nil, &InkV8Error{Source: execJs, Err: fmt.Errorf("%+v, ", JSError)}
		}
		return nil, &InkV8Error{Source: execJs, Err: err}
	}
	res, err := c.decodeV8Res(options, val)
	if err != nil {
		return nil, &InkV8Error{Source: execJs, Err: err}
	}

	return res, nil
}

func (*impl) compileFile(source string) (*v8go.CompilerCachedData, error) {
	// creates a new JavaScript VM
	iso1 := v8go.NewIsolate()
	// compile script to get cached data
	script1, err := iso1.CompileUnboundScript(source, inkVirtualJSFilePath, v8go.CompileOptions{})
	if err != nil {
		return nil, &InkV8Error{Source: source, Err: err}
	}
	return script1.CreateCodeCache(), nil
}

func (c *impl) loadFileAndCompile(
	path string,
	mutex *sync.RWMutex,
	store map[string]fileCache,
	transform func(path string, sFile string) string,
) (*fileCache, error) {
	res := func() *fileCache {
		mutex.RLock()
		defer mutex.RUnlock()
		if res, f := store[path]; f {
			return &res
		}
		return nil
	}()
	if res != nil {
		return res, nil
	}
	mutex.Lock()
	defer mutex.Unlock()
	if res, f := store[path]; f {
		return &res, nil
	}

	bFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	sFile := string(bFile)
	if transform != nil {
		sFile = transform(path, sFile)
	}

	codeCache, err := c.compileFile(sFile)
	if err != nil {
		return nil, err
	}

	fileCache := fileCache{
		content:   sFile,
		codeCache: codeCache,
	}

	store[path] = fileCache
	return &fileCache, nil
}

func (c *impl) loadFileInContext(
	iso1 *v8go.Isolate,
	inkCtx *v8go.Context,
	path string,
	mutex *sync.RWMutex,
	store map[string]fileCache,
	transform func(path string, sFile string) string,
) error {
	fc, err := c.loadFileAndCompile(path, mutex, store, transform)
	if err != nil {
		return err
	}
	// compile script in new isolate with cached data
	script, err := iso1.CompileUnboundScript(fc.content, inkVirtualJSFilePath, v8go.CompileOptions{CachedData: fc.codeCache})
	if err != nil {
		return &InkV8Error{Source: fc.content, Err: err}
	}
	//execute the script, it will be loaded into the context
	_, err = script.Run(inkCtx)
	if err != nil {
		return &InkV8Error{Source: fc.content, Err: err}
	}
	return nil
}
