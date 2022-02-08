package inkcallerv8

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/vincentkerdraon/inkcaller"
)

func Test_impl_Call_WhenStoryJS(t *testing.T) {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/ink_engine/1.0/ink.js")

	tr := NewInkCallerV8()
	//////
	seed := inkcaller.Seed(2)
	stateEncoded, err := tr.Call(ctx, engineFilePath, storyFilePath, &seed, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"choiceThreads":{"1":{"callstack":[{"cPath":"0","idx":25,"exp":false,"type":0,"temp":{}}],"threadIndex":1,"previousContentObject":"0.24"},"2":{"callstack":[{"cPath":"0","idx":32,"exp":false,"type":0,"temp":{}}],"threadIndex":2,"previousContentObject":"0.31"}},"callstackThreads":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{}}],"threadIndex":0,"previousContentObject":"0.32"}],"threadCounter":2},"variablesState":{"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesAvailable":{"list":{}},"Level":0,"DEBUG":1},"evalStack":[],"outputStream":["^It is possible to inject external data into the ink state, and read it from ink to make conditions. The variables must be declared initially in ink.","\n"],"currentChoices":[{"text":"Visit the Hub","index":0,"originalChoicePath":"0.25","originalThreadIndex":1,"targetPath":"0.c-0"},{"text":"INK_DEBUG","index":0,"originalChoicePath":"0.32","originalThreadIndex":2,"targetPath":"0.c-1"}],"visitCounts":{"":1},"turnIndices":{},"turnIdx":-1,"storySeed":2,"previousRandom":0,"inkSaveVersion":8,"inkFormatVersion":19}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS BeginStory,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	//////
	choice := inkcaller.ChoiceIndex(0)
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, nil, &choice)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"choiceThreads":{"3":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0,"temp":{}}],"threadIndex":3,"previousContentObject":"Hub.0.4"},"4":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0,"temp":{}}],"threadIndex":4,"previousContentObject":"Hub.0.23"}},"callstackThreads":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{}}],"threadIndex":1,"previousContentObject":"Hub.0.24"}],"threadCounter":4},"variablesState":{"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesAvailable":{"list":{}},"Level":0,"DEBUG":1},"evalStack":[],"outputStream":["^ "],"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":3,"targetPath":"Hub.0.c-0"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":4,"targetPath":"Hub.0.c-3"}],"visitCounts":{"":1,"0.c-0":1,"Hub":1},"turnIndices":{},"turnIdx":0,"storySeed":2,"previousRandom":0,"inkSaveVersion":8,"inkFormatVersion":19}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS First Choice,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	//////
	knotName := inkcaller.KnotName("Hub")
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, &knotName, nil)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"choiceThreads":{"5":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0,"temp":{}}],"threadIndex":5,"previousContentObject":"Hub.0.4"},"6":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0,"temp":{}}],"threadIndex":6,"previousContentObject":"Hub.0.23"}},"callstackThreads":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{}}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":6},"variablesState":{"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesAvailable":{"list":{}},"Level":0,"DEBUG":1},"evalStack":[],"outputStream":[],"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":5,"targetPath":"Hub.0.c-0"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":6,"targetPath":"Hub.0.c-3"}],"visitCounts":{"":2,"0.c-0":1,"Hub":2},"turnIndices":{},"turnIdx":1,"storySeed":2,"previousRandom":0,"inkSaveVersion":8,"inkFormatVersion":19}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS Hub,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}
}

func Test_impl_Call_WhenStoryJSON(t *testing.T) {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/ink_engine/1.0/ink.js")

	tr := NewInkCallerV8()
	_, err := tr.Call(ctx, engineFilePath, storyFilePath, nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}
