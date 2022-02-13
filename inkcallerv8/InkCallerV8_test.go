package inkcallerv8

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vincentkerdraon/inkcaller"
)

func Test_impl_Call_WhenStoryJSON(t *testing.T) {
	//This is an integration test, following story_demo.json

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

	expected := `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"0.32"}],"threadCounter":2},"outputStream":["^It is possible to inject external data into the ink state, and read it from ink to make conditions. The variables must be declared initially in ink.","\n"],"choiceThreads":{"1":{"callstack":[{"cPath":"0","idx":25,"exp":false,"type":0}],"threadIndex":1,"previousContentObject":"0.24"},"2":{"callstack":[{"cPath":"0","idx":32,"exp":false,"type":0}],"threadIndex":2,"previousContentObject":"0.31"}},"currentChoices":[{"text":"Visit the Hub","index":0,"originalChoicePath":"0.25","originalThreadIndex":1,"targetPath":"0.c-0"},{"text":"INK_DEBUG","index":0,"originalChoicePath":"0.32","originalThreadIndex":2,"targetPath":"0.c-1"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":-1,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS BeginStory,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	//////
	knotName := inkcaller.KnotName("Hub")
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, &knotName, nil)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":4},"outputStream":[],"choiceThreads":{"3":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":3,"previousContentObject":"Hub.0.4"},"4":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":4,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":3,"targetPath":"Hub.0.c-0"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":4,"targetPath":"Hub.0.c-3"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":0,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS Hub,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	//////
	choice := inkcaller.ChoiceIndex(0) //Start Scene1
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, nil, &choice)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{"$r":{"^->":"Scene1.0.2.$r1"}}}],"threadIndex":3,"previousContentObject":"Scene1.0.2.8"}],"threadCounter":5},"outputStream":["^ ","^Welcome to Scene1.","\n"],"choiceThreads":{"5":{"callstack":[{"cPath":"Scene1.0.2","idx":8,"exp":false,"type":0,"temp":{"$r":{"^->":"Scene1.0.2.$r1"}}}],"threadIndex":5,"previousContentObject":"Scene1.0.2.7"}},"currentChoices":[{"text":"Go to Scene1_1","index":0,"originalChoicePath":"Scene1.0.2.8","originalThreadIndex":5,"targetPath":"Scene1.0.c-0"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":1,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS First choice,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	//////
	choice = inkcaller.ChoiceIndex(0) //Go to Scene1_1
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, nil, &choice)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{"$r":{"^->":"Scene1.0.c-0.$r2"}}}],"threadIndex":5,"previousContentObject":"Scene1_1.13"}],"threadCounter":5},"outputStream":["^Note you can't go back to the hub from here if you are using inky editor or the web export. To mitigate that, you could for example create a \"INK_DEBUG\" knot available in all the dead-ends to test outside the lib. And then parse and remove this choice when in production.","\n"],"currentChoices":[]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}}},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":2,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS Second choice,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	knotName = inkcaller.KnotName("Hub")
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, &knotName, nil)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":8},"outputStream":[],"choiceThreads":{"6":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":6,"previousContentObject":"Hub.0.4"},"7":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":7,"previousContentObject":"Hub.0.13"},"8":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":8,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":6,"targetPath":"Hub.0.c-0"},{"text":"Start Scene2 (only visible after Scene1)","index":0,"originalChoicePath":"Hub.0.14","originalThreadIndex":7,"targetPath":"Hub.0.c-1"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":8,"targetPath":"Hub.0.c-3"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}}},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":3,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS Back to Hub,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	//////
	choice = inkcaller.ChoiceIndex(0) //Start Scene2 (only visible after Scene1)
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, nil, &choice)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{"$r":{"^->":"Scene1.0.2.$r1"}}}],"threadIndex":6,"previousContentObject":"Scene1.0.2.8"}],"threadCounter":9},"outputStream":["^ ","^Welcome to Scene1.","\n"],"choiceThreads":{"9":{"callstack":[{"cPath":"Scene1.0.2","idx":8,"exp":false,"type":0,"temp":{"$r":{"^->":"Scene1.0.2.$r1"}}}],"threadIndex":9,"previousContentObject":"Scene1.0.2.7"}},"currentChoices":[{"text":"Go to Scene1_1","index":0,"originalChoicePath":"Scene1.0.2.8","originalThreadIndex":9,"targetPath":"Scene1.0.c-0"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}}},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":4,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS Another choice,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}

	////// get access to scene 3
	stateEncodedWithGameData := inkcaller.StateEncoded(strings.ReplaceAll(string(*stateEncoded), `"variablesState":{`, `"variablesState":{"Level":1,`))
	knotName = inkcaller.KnotName("Hub")
	stateEncoded, err = tr.Call(ctx, engineFilePath, storyFilePath, nil, &stateEncodedWithGameData, &knotName, nil)
	if err != nil {
		t.Fatal(err)
	}

	expected = `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":13},"outputStream":[],"choiceThreads":{"10":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":10,"previousContentObject":"Hub.0.4"},"11":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":11,"previousContentObject":"Hub.0.13"},"12":{"callstack":[{"cPath":"Hub.0","idx":23,"exp":false,"type":0}],"threadIndex":12,"previousContentObject":"Hub.0.22"},"13":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":13,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":10,"targetPath":"Hub.0.c-0"},{"text":"Start Scene2 (only visible after Scene1)","index":0,"originalChoicePath":"Hub.0.14","originalThreadIndex":11,"targetPath":"Hub.0.c-1"},{"text":"Start Scene3 (only visible by changing ink internal state)","index":0,"originalChoicePath":"Hub.0.23","originalThreadIndex":12,"targetPath":"Hub.0.c-2"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":13,"targetPath":"Hub.0.c-3"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}},"Level":1},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":5,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`
	if string(*stateEncoded) != expected {
		t.Errorf(fmt.Sprintf("Test_impl_Call_WhenStoryJS Another choice,\ngot =%s\nwant=%s", *stateEncoded, expected))
	}
}

func Test_impl_Call_WhenStoryJS(t *testing.T) {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.js")
	engineFilePath, _ := filepath.Abs("../assets_demo/ink_engine/1.0/ink.js")

	tr := NewInkCallerV8()
	_, err := tr.Call(ctx, engineFilePath, storyFilePath, nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}
