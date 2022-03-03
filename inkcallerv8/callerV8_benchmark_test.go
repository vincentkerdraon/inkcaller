package inkcallerv8

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
	"golang.org/x/sync/errgroup"
)

// goos: linux
// goarch: amd64
// pkg: gitlab.com/eclypsaine/merchantai/inktranslator
// cpu: AMD Ryzen 9 3950X 16-Core Processor
//20220207 changes ink asset
// Benchmark_Translator-32    	     274	   3955527 ns/op	    2978 B/op	      34 allocs/op
//20220210 using inkjs@2.0.0
// Benchmark_Translator-32    	     294	   3621818 ns/op	    2900 B/op	      34 allocs/op
// Benchmark_Translator-32    	     302	   3585286 ns/op	    2860 B/op	      34 allocs/op
// Benchmark_Translator-32    	     303	   3633023 ns/op	    2856 B/op	      34 allocs/op
//20220301 api completely changed. Still inkjs@2.0.0. Now using ink api instead of reading state. Decoding types in V8.
//When no output.
// Benchmark_Translator-32    	     318	   3331378 ns/op	    2113 B/op	      36 allocs/op
// Benchmark_Translator-32    	     325	   3219831 ns/op	    2081 B/op	      36 allocs/op
// Benchmark_Translator-32    	     330	   3339780 ns/op	    2074 B/op	      36 allocs/op
//When output StateOut
// Benchmark_Translator-32    	     313	   3575697 ns/op	    4491 B/op	      47 allocs/op
// Benchmark_Translator-32    	     296	   3660793 ns/op	    4556 B/op	      47 allocs/op
//When output Lines + LineTags + GlobalTags + Choices
// Benchmark_Translator-32    	     330	   3591482 ns/op	    9229 B/op	     343 allocs/op
// Benchmark_Translator-32    	     326	   3394489 ns/op	    9230 B/op	     343 allocs/op
//When output StateOut + Lines + LineTags + GlobalTags + Choices
// Benchmark_Translator-32    	     282	   3758420 ns/op	   12073 B/op	     354 allocs/op
// Benchmark_Translator-32    	     283	   3778318 ns/op	   12087 B/op	     354 allocs/op
func Benchmark_Translator(b *testing.B) {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/inkjs_engine/2.0.0/ink.js")

	c := NewInkCallerV8()
	for n := 0; n < b.N; n++ {
		_, err := c.Call(ctx, engineFilePath, storyFilePath,
			inkcallerlib.WithOutputChoices(true),
			inkcallerlib.WithOutputLines(true),
			inkcallerlib.WithOutputLineTags(true),
			inkcallerlib.WithOutputStateOut(true),
			inkcallerlib.WithOutputGlobalTags(true),
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}

// goos: linux
// goarch: amd64
// pkg: gitlab.com/eclypsaine/merchantai/inktranslator
// cpu: AMD Ryzen 9 3950X 16-Core Processor
//20220207 changes ink asset / routines := 10
// Benchmark_Translator_race-32    	     188	   6639735 ns/op	    3457 B/op	      33 allocs/op
//20220210 using inkjs@2.0.0 / routines := 10
// Benchmark_Translator_race-32    	     192	   5865029 ns/op	    3564 B/op	      34 allocs/op
//20220210 using inkjs@2.0.0 / routines := 100
// Benchmark_Translator_race-32    	     327	   5908055 ns/op	    2746 B/op	      32 allocs/op
//20220210 using inkjs@2.0.0 / routines := 1
// Benchmark_Translator_race-32    	     307	   3618749 ns/op	    2861 B/op	      34 allocs/op
//20220301 api completely changed. Still inkjs@2.0.0. Now using ink api instead of reading state
//20220301 routines := 1
// Benchmark_Translator_race-32    	     276	   3721015 ns/op	    9673 B/op	     208 allocs/op
//20220301 routines := 10
// Benchmark_Translator_race-32    	     166	   6719884 ns/op	   10248 B/op	     201 allocs/op
// Benchmark_Translator_race-32    	     170	   6949299 ns/op	   10513 B/op	     208 allocs/op
//20220301 routines := 100
// Benchmark_Translator_race-32    	     277	   5094197 ns/op	    7514 B/op	     151 allocs/op
func Benchmark_Translator_race(b *testing.B) {
	routines := 10

	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/inkjs_engine/2.0.0/ink.js")
	ctx := context.Background()
	c := NewInkCallerV8()
	eg, _ := errgroup.WithContext(context.Background())
	for r := routines; r > 0; r-- {
		eg.Go(func() error {
			for n := b.N / routines; n > 0; n-- {
				_, err := c.Call(ctx, engineFilePath, storyFilePath)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	eg.Wait()
}
