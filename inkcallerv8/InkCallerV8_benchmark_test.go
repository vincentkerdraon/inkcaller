package inkcallerv8

import (
	"context"
	"path/filepath"
	"testing"

	"golang.org/x/sync/errgroup"
)

// goos: linux
// goarch: amd64
// pkg: gitlab.com/eclypsaine/merchantai/inktranslator
// cpu: AMD Ryzen 9 3950X 16-Core Processor
//20220207 changes ink asset
// Benchmark_Translator-32    	     274	   3955527 ns/op	    2978 B/op	      34 allocs/op
func Benchmark_Translator(b *testing.B) {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/ink_engine/1.0/ink.js")

	c := NewInkCallerV8()
	for n := 0; n < b.N; n++ {
		_, err := c.Call(ctx, engineFilePath, storyFilePath, nil, nil, nil, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// goos: linux
// goarch: amd64
// pkg: gitlab.com/eclypsaine/merchantai/inktranslator
// cpu: AMD Ryzen 9 3950X 16-Core Processor
//20220207 changes ink asset
// Benchmark_Translator_race-32    	     188	   6639735 ns/op	    3457 B/op	      33 allocs/op
func Benchmark_Translator_race(b *testing.B) {
	routines := 10

	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/ink_engine/1.0/ink.js")
	ctx := context.Background()
	c := NewInkCallerV8()
	eg, _ := errgroup.WithContext(context.Background())
	for r := routines; r > 0; r-- {
		eg.Go(func() error {
			for n := b.N / routines; n > 0; n-- {
				_, err := c.Call(ctx, engineFilePath, storyFilePath, nil, nil, nil, nil)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	eg.Wait()
}
