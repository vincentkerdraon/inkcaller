package inkcaller

import (
	"encoding/json"
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
			args: args{sEncoded: `{"choiceThreads":{"4":{"callstack":[{"cPath":"SCENE_Tuto1rst.0.6","idx":8,"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.6.$r1"}}}],"threadIndex":4,"previousContentObject":"SCENE_Tuto1rst.0.6.7"},"5":{"callstack":[{"cPath":"SCENE_Tuto1rst.0.7","idx":8,"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.7.$r1"}}}],"threadIndex":5,"previousContentObject":"SCENE_Tuto1rst.0.7.7"}},"callstackThreads":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.7.$r1"}}}],"threadIndex":1,"previousContentObject":"SCENE_Tuto1rst.0.7.8"}],"threadCounter":5},"variablesState":{"CharactersID":{"list":{},"origins":["CharactersID"]},"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesPossible":{"list":{"ScenesID.Tuto1":2,"ScenesID.TutoTemps":4}},"PresentInShop":{"list":{}},"CurrentTime":15,"CurrencyMain":50000,"TutoTempsRetourGuerrier":0},"evalStack":[],"outputStream":["^ ","^Ceci est le tuto 1. Une première quete super facile !","\n"],"currentChoices":[{"text":"Veux tu choisir le premier choix?","index":0,"originalChoicePath":"SCENE_Tuto1rst.0.6.8","originalThreadIndex":4,"targetPath":"SCENE_Tuto1rst.0.c-0"},{"text":"Ou le deuxième?","index":1,"originalChoicePath":"SCENE_Tuto1rst.0.7.8","originalThreadIndex":5,"targetPath":"SCENE_Tuto1rst.0.c-1"}],"visitCounts":{"":1,"MENU.NewGame":1,"MENU":1,"Smalltalk":1,"Smalltalk.0":1,"Hub":1,"Hub.0.c-0":1,"SCENE_Tuto1rst":1},"turnIndices":{},"turnIdx":0,"storySeed":30,"previousRandom":0,"inkSaveVersion":8,"inkFormatVersion":19}`},
			want: &InkState{
				CallstackThreads: CallstackThreads{
					Threads: []Thread{
						{Callstack: []ThreadCallstack{
							{Temp: Temp{R: R{Empty: "SCENE_Tuto1rst.0.7.$r1"}}},
						}},
					},
				},
				VariablesState: json.RawMessage(`{"CharactersID":{"list":{},"origins":["CharactersID"]},"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesPossible":{"list":{"ScenesID.Tuto1":2,"ScenesID.TutoTemps":4}},"PresentInShop":{"list":{}},"CurrentTime":15,"CurrencyMain":50000,"TutoTempsRetourGuerrier":0}`),
				OutputStream:   []string{"^ ", "^Ceci est le tuto 1. Une première quete super facile !", "\n"},
				CurrentChoices: []CurrentChoice{
					{Index: 0, Text: "Veux tu choisir le premier choix?", TargetPath: "SCENE_Tuto1rst.0.c-0"},
					{Index: 1, Text: "Ou le deuxième?", TargetPath: "SCENE_Tuto1rst.0.c-1"},
				},
				VisitCounts: map[string]int64{
					"":               1,
					"Hub":            1,
					"Hub.0.c-0":      1,
					"MENU":           1,
					"MENU.NewGame":   1,
					"SCENE_Tuto1rst": 1,
					"Smalltalk":      1,
					"Smalltalk.0":    1,
				},
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
			sEncoded:     `{"choiceThreads":{"4":{"callstack":[{"cPath":"SCENE_Tuto1rst.0.6","idx":8,"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.6.$r1"}}}],"threadIndex":4,"previousContentObject":"SCENE_Tuto1rst.0.6.7"},"5":{"callstack":[{"cPath":"SCENE_Tuto1rst.0.7","idx":8,"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.7.$r1"}}}],"threadIndex":5,"previousContentObject":"SCENE_Tuto1rst.0.7.7"}},"callstackThreads":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.7.$r1"}}}],"threadIndex":1,"previousContentObject":"SCENE_Tuto1rst.0.7.8"}],"threadCounter":5},"variablesState":{"CharactersID":{"list":{},"origins":["CharactersID"]},"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesPossible":{"list":{"ScenesID.Tuto1":2,"ScenesID.TutoTemps":4}},"PresentInShop":{"list":{}},"CurrentTime":15,"CurrencyMain":50000,"TutoTempsRetourGuerrier":0},"evalStack":[],"outputStream":["^ ","^Ceci est le tuto 1. Une première quete super facile !","\n"],"currentChoices":[{"text":"Veux tu choisir le premier choix?","index":0,"originalChoicePath":"SCENE_Tuto1rst.0.6.8","originalThreadIndex":4,"targetPath":"SCENE_Tuto1rst.0.c-0"},{"text":"Ou le deuxième?","index":1,"originalChoicePath":"SCENE_Tuto1rst.0.7.8","originalThreadIndex":5,"targetPath":"SCENE_Tuto1rst.0.c-1"}],"visitCounts":{"":1,"MENU.NewGame":1,"MENU":1,"Smalltalk":1,"Smalltalk.0":1,"Hub":1,"Hub.0.c-0":1,"SCENE_Tuto1rst":1},"turnIndices":{},"turnIdx":0,"storySeed":30,"previousRandom":0,"inkSaveVersion":8,"inkFormatVersion":19}`,
			args:         args{gameModelV: struct{ CurrentTime int }{CurrentTime: 30}},
			wantContains: []string{`"CurrentTime":30`, `"CurrencyMain":50000`, `"storySeed":30`},
		},
		{
			name:     "when game model list",
			sEncoded: `{"choiceThreads":{"4":{"callstack":[{"cPath":"SCENE_Tuto1rst.0.6","idx":8,"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.6.$r1"}}}],"threadIndex":4,"previousContentObject":"SCENE_Tuto1rst.0.6.7"},"5":{"callstack":[{"cPath":"SCENE_Tuto1rst.0.7","idx":8,"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.7.$r1"}}}],"threadIndex":5,"previousContentObject":"SCENE_Tuto1rst.0.7.7"}},"callstackThreads":{"threads":[{"callstack":[{"exp":false,"type":0,"temp":{"temp_money":0,"$r":{"^->":"SCENE_Tuto1rst.0.7.$r1"}}}],"threadIndex":1,"previousContentObject":"SCENE_Tuto1rst.0.7.8"}],"threadCounter":5},"variablesState":{"CharactersID":{"list":{},"origins":["CharactersID"]},"ScenesID":{"list":{},"origins":["ScenesID"]},"ScenesPossible":{"list":{"ScenesID.Tuto1":2,"ScenesID.TutoTemps":4}},"PresentInShop":{"list":{}},"CurrentTime":15,"CurrencyMain":50000,"TutoTempsRetourGuerrier":0},"evalStack":[],"outputStream":["^ ","^Ceci est le tuto 1. Une première quete super facile !","\n"],"currentChoices":[{"text":"Veux tu choisir le premier choix?","index":0,"originalChoicePath":"SCENE_Tuto1rst.0.6.8","originalThreadIndex":4,"targetPath":"SCENE_Tuto1rst.0.c-0"},{"text":"Ou le deuxième?","index":1,"originalChoicePath":"SCENE_Tuto1rst.0.7.8","originalThreadIndex":5,"targetPath":"SCENE_Tuto1rst.0.c-1"}],"visitCounts":{"":1,"MENU.NewGame":1,"MENU":1,"Smalltalk":1,"Smalltalk.0":1,"Hub":1,"Hub.0.c-0":1,"SCENE_Tuto1rst":1},"turnIndices":{},"turnIdx":0,"storySeed":30,"previousRandom":0,"inkSaveVersion":8,"inkFormatVersion":19}`,
			args: args{gameModelV: struct {
				CurrentTime   int
				PresentInShop struct {
					List map[string]int `json:"list"`
				}
			}{
				CurrentTime: 30,
				PresentInShop: struct {
					List map[string]int `json:"list"`
				}{
					map[string]int{
						"CharactersID.Michel": 2,
					},
				},
			}},
			wantContains: []string{`"CurrentTime":30`, `"CurrencyMain":50000`, `"storySeed":30`, `"PresentInShop":{"list":{"CharactersID.Michel":2}}`},
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
