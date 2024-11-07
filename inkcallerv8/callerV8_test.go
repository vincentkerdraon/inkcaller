package inkcallerv8

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
)

func Test_impl_Call_WhenStoryJSON(t *testing.T) {
	//This is an integration test, following story_demo.json

	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/inkjs_engine/2.3.0/ink.js")
	checkFunc := callAndValidate(context.Background(), t, NewInkCallerV8(), storyFilePath, engineFilePath)

	//////

	expectedStateOut := inkcallerlib.StateEncoded(`{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"0.38"}],"threadCounter":2},"outputStream":["^It is possible to inject external data into the ink state, and read it from ink to make conditions. The variables must be declared initially in ink.","\n"],"choiceThreads":{"1":{"callstack":[{"cPath":"0","idx":31,"exp":false,"type":0}],"threadIndex":1,"previousContentObject":"0.30"},"2":{"callstack":[{"cPath":"0","idx":38,"exp":false,"type":0}],"threadIndex":2,"previousContentObject":"0.37"}},"currentChoices":[{"text":"Visit the Hub","index":0,"originalChoicePath":"0.31","originalThreadIndex":1,"targetPath":"0.c-0","tags":[]},{"text":"INK_DEBUG","index":0,"originalChoicePath":"0.38","originalThreadIndex":2,"targetPath":"0.c-1","tags":[]}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":-1,"storySeed":2,"previousRandom":0,"inkSaveVersion":10,"inkFormatVersion":21}`)
	expected := inkcallerlib.InkCallerOutput{
		StateOut:   &expectedStateOut,
		GlobalTags: &[]inkcallerlib.Tag{"tag_global1", "tag_global2"},
	}
	out := checkFunc("start", expected,
		inkcallerlib.WithInputSeed(inkcallerlib.Seed(2)),
		inkcallerlib.WithOutputGlobalTags(true),
		inkcallerlib.WithOutputStateOut(true),
	)

	//////
	expectedStateOut = inkcallerlib.StateEncoded(`{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":4},"outputStream":[],"choiceThreads":{"3":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":3,"previousContentObject":"Hub.0.4"},"4":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":4,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":3,"targetPath":"Hub.0.c-0","tags":[]},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":4,"targetPath":"Hub.0.c-3","tags":[]}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":0,"storySeed":2,"previousRandom":0,"inkSaveVersion":10,"inkFormatVersion":21}`)
	turnIndex := inkcallerlib.TurnIndex(0)
	expected = inkcallerlib.InkCallerOutput{
		StateOut: &expectedStateOut,
		Choices: &[]inkcallerlib.Choice{
			{Index: 0, Text: "Start Scene1", SourcePath: "Hub.0.5"},
		},
		TurnIndex: &turnIndex,
	}
	out = checkFunc("HUB 1", expected,
		inkcallerlib.WithInputStateIn(*out.StateOut),
		inkcallerlib.WithInputKnotName(inkcallerlib.KnotName("Hub")),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputGlobalTags(false),
		inkcallerlib.WithOutputLines(false),
		inkcallerlib.WithOutputChoices(true),
		inkcallerlib.WithOutputTurnIndex(true),
	)

	//////
	expectedStateOut = inkcallerlib.StateEncoded(`{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":3,"previousContentObject":"Scene1.0.13"}],"threadCounter":5},"outputStream":["^and with parenthesis) ","^Welcome to Scene1. ","#","^tag1 ","/#","#","^tag2","/#","\n"],"choiceThreads":{"5":{"callstack":[{"cPath":"Scene1.0","idx":13,"exp":false,"type":0}],"threadIndex":5,"previousContentObject":"Scene1.0.12"}},"currentChoices":[{"text":"Go to Scene1_1","index":0,"originalChoicePath":"Scene1.0.13","originalThreadIndex":5,"targetPath":"Scene1.0.c-0","tags":[]}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":1,"storySeed":2,"previousRandom":0,"inkSaveVersion":10,"inkFormatVersion":21}`)
	turnIndex = inkcallerlib.TurnIndex(1)
	expected = inkcallerlib.InkCallerOutput{
		StateOut: &expectedStateOut,
		Lines: &[]inkcallerlib.Line{
			//The "On multiple..." comes from the choice.
			//The first line of the knot is "Welcome to Scene1"
			//and contains the tags
			{Text: "(On multiple lines\n", Tags: &[]inkcallerlib.Tag{}},
			{Text: "and with parenthesis) Welcome to Scene1.\n", Tags: &[]inkcallerlib.Tag{"tag1", "tag2"}},
		},
		Choices: &[]inkcallerlib.Choice{
			{Index: 0, Text: "Go to Scene1_1", SourcePath: "Scene1.0.13"},
		},
		TurnIndex: &turnIndex,
	}
	out = checkFunc("Start Scene1", expected,
		inkcallerlib.WithInputStateIn(*out.StateOut),
		inkcallerlib.WithInputChoiceIndex(inkcallerlib.ChoiceIndex(0)),
		inkcallerlib.WithOutputGlobalTags(false),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputLines(true),
		inkcallerlib.WithOutputChoices(true),
		inkcallerlib.WithOutputLineTags(true),
		inkcallerlib.WithOutputTurnIndex(true),
	)

	//////
	expectedStateOut = inkcallerlib.StateEncoded(`{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":5,"previousContentObject":"Scene1_1.13"}],"threadCounter":5},"outputStream":["^Note you can't go back to the hub from here if you are using inky editor or the web export. To mitigate that, you could for example create a \"INK_DEBUG\" knot available in all the dead-ends to test outside the lib. And then parse and remove this choice when in production.","\n"],"currentChoices":[]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}}},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":2,"storySeed":2,"previousRandom":0,"inkSaveVersion":10,"inkFormatVersion":21}`)
	expected = inkcallerlib.InkCallerOutput{
		StateOut: &expectedStateOut,
		Lines: &[]inkcallerlib.Line{
			{Text: "Welcome to Scene1_1.\n", Tags: &[]inkcallerlib.Tag{}},
			{Text: "That's the end of this scene. (no choice available)\n", Tags: &[]inkcallerlib.Tag{}},
			{Text: "Now Scene2 is available from the hub!\n", Tags: &[]inkcallerlib.Tag{}},
			{Text: `Note you can't go back to the hub from here if you are using inky editor or the web export. To mitigate that, you could for example create a "INK_DEBUG" knot available in all the dead-ends to test outside the lib. And then parse and remove this choice when in production.` + "\n", Tags: &[]inkcallerlib.Tag{}},
		},
		Choices: &[]inkcallerlib.Choice{},
	}
	out = checkFunc("Continue to Scene1_1", expected,
		inkcallerlib.WithInputStateIn(*out.StateOut),
		inkcallerlib.WithInputChoiceIndex(inkcallerlib.ChoiceIndex(0)),
		inkcallerlib.WithOutputGlobalTags(false),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputLines(true),
		inkcallerlib.WithOutputChoices(true),
		inkcallerlib.WithOutputLineTags(true),
	)

	//////
	expectedStateOut = inkcallerlib.StateEncoded(`{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":8},"outputStream":[],"choiceThreads":{"6":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":6,"previousContentObject":"Hub.0.4"},"7":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":7,"previousContentObject":"Hub.0.13"},"8":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":8,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":6,"targetPath":"Hub.0.c-0","tags":[]},{"text":"Start Scene2 (only visible after Scene1)","index":1,"originalChoicePath":"Hub.0.14","originalThreadIndex":7,"targetPath":"Hub.0.c-1","tags":[]},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":8,"targetPath":"Hub.0.c-3","tags":[]}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}}},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":3,"storySeed":2,"previousRandom":0,"inkSaveVersion":10,"inkFormatVersion":21}`)
	expected = inkcallerlib.InkCallerOutput{
		StateOut: &expectedStateOut,
		Choices: &[]inkcallerlib.Choice{
			{Index: 0, Text: "Start Scene1", SourcePath: "Hub.0.5"},
			{Index: 1, Text: "Start Scene2 (only visible after Scene1)", SourcePath: "Hub.0.14"},
		},
	}
	out = checkFunc("HUB 2", expected,
		inkcallerlib.WithInputStateIn(*out.StateOut),
		inkcallerlib.WithInputKnotName(inkcallerlib.KnotName("Hub")),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputGlobalTags(false),
		inkcallerlib.WithOutputLines(false),
		inkcallerlib.WithOutputChoices(true),
	)

	//////
	// get access to scene 3, not possible from inside ink, needs a change in variables.
	stateEncodedWithGameData := inkcallerlib.StateEncoded(strings.ReplaceAll(string(*out.StateOut), `"variablesState":{`, `"variablesState":{"Level":1,`))

	expectedStateOut = inkcallerlib.StateEncoded(`{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":12},"outputStream":[],"choiceThreads":{"9":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":9,"previousContentObject":"Hub.0.4"},"10":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":10,"previousContentObject":"Hub.0.13"},"11":{"callstack":[{"cPath":"Hub.0","idx":23,"exp":false,"type":0}],"threadIndex":11,"previousContentObject":"Hub.0.22"},"12":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":12,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":9,"targetPath":"Hub.0.c-0","tags":[]},{"text":"Start Scene2 (only visible after Scene1)","index":1,"originalChoicePath":"Hub.0.14","originalThreadIndex":10,"targetPath":"Hub.0.c-1","tags":[]},{"text":"Start Scene3 (only visible by changing ink internal state)","index":2,"originalChoicePath":"Hub.0.23","originalThreadIndex":11,"targetPath":"Hub.0.c-2","tags":[]},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":12,"targetPath":"Hub.0.c-3","tags":[]}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}},"Level":1},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":4,"storySeed":2,"previousRandom":0,"inkSaveVersion":10,"inkFormatVersion":21}`)
	expected = inkcallerlib.InkCallerOutput{
		StateOut: &expectedStateOut,
		Choices: &[]inkcallerlib.Choice{
			{Index: 0, Text: "Start Scene1", SourcePath: "Hub.0.5"},
			{Index: 1, Text: "Start Scene2 (only visible after Scene1)", SourcePath: "Hub.0.14"},
			{Index: 2, Text: "Start Scene3 (only visible by changing ink internal state)", SourcePath: "Hub.0.23"},
		},
	}
	out = checkFunc("HUB 3", expected,
		inkcallerlib.WithInputStateIn(stateEncodedWithGameData),
		inkcallerlib.WithInputKnotName(inkcallerlib.KnotName("Hub")),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputGlobalTags(false),
		inkcallerlib.WithOutputLines(false),
		inkcallerlib.WithOutputChoices(true),
	)
	_ = out
}

func callAndValidate(ctx context.Context, t *testing.T, tr *impl, storyFilePath string, engineFilePath string) func(testCase string, expected inkcallerlib.InkCallerOutput, opts ...inkcallerlib.InkCallerOptionsFunc) *inkcallerlib.InkCallerOutput {
	return func(testCase string, expected inkcallerlib.InkCallerOutput, opts ...inkcallerlib.InkCallerOptionsFunc) *inkcallerlib.InkCallerOutput {
		out, err := tr.Call(ctx, engineFilePath, storyFilePath, opts...)
		if err != nil {
			t.Fatal(err)
		}

		//this struct is a pain to test!

		if (out.TurnIndex == nil || expected.TurnIndex == nil) && expected.TurnIndex != out.TurnIndex {
			t.Fatalf(fmt.Sprintf("when %q, a TurnIndex is nil\ngot =%+v\nwant=%+v", testCase, out, expected))
		}
		if expected.TurnIndex != nil && *out.TurnIndex != *expected.TurnIndex {
			t.Errorf(fmt.Sprintf("when %q,\ngot =%d\nwant=%d", testCase, *out.TurnIndex, *expected.TurnIndex))
		}

		if (out.StateOut == nil || expected.StateOut == nil) && expected.StateOut != out.StateOut {
			t.Fatalf(fmt.Sprintf("when %q, a StateOut is nil\ngot =%+v\nwant=%+v", testCase, out, expected))
		}
		if expected.StateOut != nil && *out.StateOut != *expected.StateOut {
			t.Errorf(fmt.Sprintf("when %q,\ngot =%s\nwant=%s", testCase, *out.StateOut, *expected.StateOut))
		}

		if (out.Choices == nil || expected.Choices == nil) && expected.Choices != out.Choices {
			t.Fatalf(fmt.Sprintf("when %q, a choice is nil,\ngot =%+v\nwant=%+v", testCase, out, expected))
		}
		if expected.Choices != nil {
			if len(*expected.Choices) != len(*out.Choices) {
				t.Fatalf(fmt.Sprintf("when %q,\ngot =%+v\nwant=%+v", testCase, *out.Choices, *expected.Choices))
			}
			for i, c := range *out.Choices {
				if (*expected.Choices)[i].Index != c.Index || (*expected.Choices)[i].Text != c.Text || (*expected.Choices)[i].SourcePath != c.SourcePath {
					t.Errorf(fmt.Sprintf("when %q, choice index %d\ngot =%+v\nwant=%+v", testCase, i, *out.Choices, *expected.Choices))
				}
			}
		}

		if (out.GlobalTags == nil || expected.GlobalTags == nil) && expected.GlobalTags != out.GlobalTags {
			t.Fatalf(fmt.Sprintf("when %q, a GlobalTags is nil,\ngot =%+v\nwant=%+v", testCase, out, expected))
		}
		if expected.GlobalTags != nil {
			if len(*expected.GlobalTags) != len(*out.GlobalTags) {
				t.Fatalf(fmt.Sprintf("when %q,\ngot =%s\nwant=%s", testCase, *out.GlobalTags, *expected.GlobalTags))
			}
			for i, c := range *out.GlobalTags {
				if (*expected.GlobalTags)[i] != c {
					t.Errorf(fmt.Sprintf("when %q,\ngot =%s\nwant=%s", testCase, *out.GlobalTags, *expected.GlobalTags))
				}
			}
		}

		if (out.Lines == nil || expected.Lines == nil) && expected.Lines != out.Lines {
			t.Fatalf(fmt.Sprintf("when %q, a Lines is nil,\ngot =%+v\nwant=%+v", testCase, out, expected))
		}
		if expected.Lines != nil {
			if len(*expected.Lines) != len(*out.Lines) {
				t.Fatalf(fmt.Sprintf("when %q,\ngot =%+v\nwant=%+v", testCase, *out.Lines, *expected.Lines))
			}
			for i, l := range *out.Lines {
				if (*expected.Lines)[i].Text != l.Text {
					t.Errorf(fmt.Sprintf("when %q, Line index %d\ngot =%+v\nwant=%+v\ngot =%q\nwant=%q",
						testCase, i, *out.Lines, *expected.Lines, l.Text, (*expected.Lines)[i].Text))
				}
				if ((*expected.Lines)[i].Tags == nil || l.Tags == nil) && (*expected.Lines)[i].Tags != l.Tags {
					t.Fatalf(fmt.Sprintf("when %q, Line index %d, Line Tags is nil,\ngot =%+v\nwant=%+v", testCase, i, *out.Lines, *expected.Lines))
				}
				if (*expected.Lines)[i].Tags != nil {
					if len(*(*expected.Lines)[i].Tags) != len(*l.Tags) {
						t.Fatalf(fmt.Sprintf("when %q, Line index %d, Line Tag is nil\ngot =%+v\nwant=%+v", testCase, i, *l.Tags, *(*expected.Lines)[i].Tags))
					}
					for j, tag := range *l.Tags {
						if tag != (*(*expected.Lines)[i].Tags)[j] {
							t.Fatalf(fmt.Sprintf("when %q, Line index %d, Line Tags invalid,\ngot =%s\nwant=%s", testCase, i, tag, *(*expected.Lines)[i].Tags))
						}
					}
				}
			}
		}
		return out
	}
}

func Test_impl_Call_WhenStoryJS(t *testing.T) {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.js")
	engineFilePath, _ := filepath.Abs("../assets_demo/inkjs_engine/2.3.0/ink.js")

	tr := NewInkCallerV8()
	_, err := tr.Call(ctx, engineFilePath, storyFilePath)
	if err != nil {
		t.Fatal(err)
	}
}
