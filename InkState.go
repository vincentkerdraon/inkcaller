package inkcaller

import (
	"encoding/json"
)

//InkState is based on Ink definition.
//We are ignoring some fields.
//This is only for reading. Writing variables into the states uses another way.
type InkState struct {
	CallstackThreads CallstackThreads `json:"callstackThreads"`
	//VariablesState contains the custom game state values, and also internal ink values. (merge, don't override)
	VariablesState json.RawMessage  `json:"variablesState"`
	OutputStream   []string         `json:"outputStream"`
	CurrentChoices []CurrentChoice  `json:"currentChoices"`
	VisitCounts    map[string]int64 `json:"visitCounts"`
}

type CallstackThreads struct {
	Threads []Thread `json:"threads"`
}

type Thread struct {
	Callstack []ThreadCallstack `json:"callstack"`
}

type ThreadCallstack struct {
	Temp Temp `json:"temp"`
}

type Temp struct {
	R R `json:"$r"`
}

type R struct {
	Empty string `json:"^->"`
}

type CurrentChoice struct {
	Text       string      `json:"text"`
	Index      ChoiceIndex `json:"index"`
	TargetPath string      `json:"targetPath"`
}

func (s InkState) ID() string {
	return s.CallstackThreads.Threads[0].Callstack[0].Temp.R.Empty
}

func NewInkStateMock(id string, text []string, choices []CurrentChoice) *InkState {
	return &InkState{
		CallstackThreads: CallstackThreads{
			Threads: []Thread{
				{Callstack: []ThreadCallstack{
					{Temp: Temp{R: R{Empty: id}}},
				}},
			},
		},
		CurrentChoices: choices,
		OutputStream:   text,
	}
}
