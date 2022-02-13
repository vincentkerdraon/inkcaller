package inkcaller

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeInkState(t *testing.T) {
	type args struct {
		sEncoded StateEncoded
	}
	tests := []struct {
		name    string
		args    args
		want    *InkState
		wantErr bool
	}{
		{
			name: "When OK",
			args: args{sEncoded: `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":13},"outputStream":[],"choiceThreads":{"10":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":10,"previousContentObject":"Hub.0.4"},"11":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":11,"previousContentObject":"Hub.0.13"},"12":{"callstack":[{"cPath":"Hub.0","idx":23,"exp":false,"type":0}],"threadIndex":12,"previousContentObject":"Hub.0.22"},"13":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":13,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":10,"targetPath":"Hub.0.c-0"},{"text":"Start Scene2 (only visible after Scene1)","index":0,"originalChoicePath":"Hub.0.14","originalThreadIndex":11,"targetPath":"Hub.0.c-1"},{"text":"Start Scene3 (only visible by changing ink internal state)","index":0,"originalChoicePath":"Hub.0.23","originalThreadIndex":12,"targetPath":"Hub.0.c-2"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":13,"targetPath":"Hub.0.c-3"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}},"Level":1},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":5,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`},
			want: &InkState{
				Flows: Flows{
					DefaultFlow: DefaultFlow{
						OutputStream: []string{},
						CurrentChoices: []CurrentChoice{
							{
								Text:       "Start Scene1",
								Index:      0,
								TargetPath: "Hub.0.c-0",
							},
							{
								Text:       "Start Scene2 (only visible after Scene1)",
								Index:      1,
								TargetPath: "Hub.0.c-1",
							},
							{
								Text:       "Start Scene3 (only visible by changing ink internal state)",
								Index:      2,
								TargetPath: "Hub.0.c-2",
							},
							{
								Text:       "",
								Index:      3,
								TargetPath: "Hub.0.c-3",
							},
						},
					},
				},
				TurnIdx: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.sEncoded.DecodeInkState()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeInkState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInkState()\ngot =%#v\nwant=%#v", got, tt.want)
			}
		})
	}
}

func TestStateEncoded_includeGameData(t *testing.T) {
	type args struct {
		gameModelV interface{}
	}
	tests := []struct {
		name         string
		sEncoded     StateEncoded
		args         args
		wantErr      bool
		wantContains []string
	}{
		{
			name:         "ok",
			sEncoded:     `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":13},"outputStream":[],"choiceThreads":{"10":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":10,"previousContentObject":"Hub.0.4"},"11":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":11,"previousContentObject":"Hub.0.13"},"12":{"callstack":[{"cPath":"Hub.0","idx":23,"exp":false,"type":0}],"threadIndex":12,"previousContentObject":"Hub.0.22"},"13":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":13,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":10,"targetPath":"Hub.0.c-0"},{"text":"Start Scene2 (only visible after Scene1)","index":0,"originalChoicePath":"Hub.0.14","originalThreadIndex":11,"targetPath":"Hub.0.c-1"},{"text":"Start Scene3 (only visible by changing ink internal state)","index":0,"originalChoicePath":"Hub.0.23","originalThreadIndex":12,"targetPath":"Hub.0.c-2"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":13,"targetPath":"Hub.0.c-3"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}},"Level":1},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":5,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`,
			args:         args{gameModelV: struct{ CurrentTime int }{CurrentTime: 30}},
			wantContains: []string{`"storySeed":2`, `"Level":1`},
		},
		{
			name:     "when game model list",
			sEncoded: `{"flows":{"DEFAULT_FLOW":{"callstack":{"threads":[{"callstack":[{"exp":false,"type":0}],"threadIndex":0,"previousContentObject":"Hub.0.24"}],"threadCounter":13},"outputStream":[],"choiceThreads":{"10":{"callstack":[{"cPath":"Hub.0","idx":5,"exp":false,"type":0}],"threadIndex":10,"previousContentObject":"Hub.0.4"},"11":{"callstack":[{"cPath":"Hub.0","idx":14,"exp":false,"type":0}],"threadIndex":11,"previousContentObject":"Hub.0.13"},"12":{"callstack":[{"cPath":"Hub.0","idx":23,"exp":false,"type":0}],"threadIndex":12,"previousContentObject":"Hub.0.22"},"13":{"callstack":[{"cPath":"Hub.0","idx":24,"exp":false,"type":0}],"threadIndex":13,"previousContentObject":"Hub.0.23"}},"currentChoices":[{"text":"Start Scene1","index":0,"originalChoicePath":"Hub.0.5","originalThreadIndex":10,"targetPath":"Hub.0.c-0"},{"text":"Start Scene2 (only visible after Scene1)","index":0,"originalChoicePath":"Hub.0.14","originalThreadIndex":11,"targetPath":"Hub.0.c-1"},{"text":"Start Scene3 (only visible by changing ink internal state)","index":0,"originalChoicePath":"Hub.0.23","originalThreadIndex":12,"targetPath":"Hub.0.c-2"},{"text":"","index":0,"originalChoicePath":"Hub.0.24","originalThreadIndex":13,"targetPath":"Hub.0.c-3"}]}},"currentFlowName":"DEFAULT_FLOW","variablesState":{"ScenesAvailable":{"list":{"ScenesID.SceneAvailable_2":1}},"Level":1,"PresentInShop":{"list":{}}},"evalStack":[],"visitCounts":{},"turnIndices":{},"turnIdx":5,"storySeed":2,"previousRandom":0,"inkSaveVersion":9,"inkFormatVersion":20}`,

			args: args{gameModelV: struct {
				//Already in state
				Level int
				//new
				Score int
				//Already in state
				PresentInShop struct {
					List map[string]int `json:"list"`
				}
				//new
				Inventory struct {
					List map[string]int `json:"list"`
				}
			}{
				Level: 2,
				Score: 10,
				PresentInShop: struct {
					List map[string]int `json:"list"`
				}{map[string]int{"CharactersID.Michel": 2}},
				Inventory: struct {
					List map[string]int "json:\"list\""
				}{map[string]int{"Bow": 1}},
			}},
			wantContains: []string{`"storySeed":2`, `"Level":2`, `"Score":10`, `"PresentInShop":{"list":{"CharactersID.Michel":2}}`, `"Inventory":{"list":{"Bow":1}}`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gd, err := tt.sEncoded.ConvertVtoMV(tt.args.gameModelV)
			if err != nil {
				t.Errorf("StateEncoded.includeGameData() error = %v", err)
			}
			got, err := tt.sEncoded.IncludeGameData(gd)
			if (err != nil) != tt.wantErr {
				t.Errorf("StateEncoded.includeGameData() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, want := range tt.wantContains {
				if !strings.Contains(string(*got), want) {
					t.Errorf(fmt.Sprintf("TestInkState_includeGameData, want contains\ngot =%s\nwant=%s", *got, want))
				}
			}
		})
	}
}
