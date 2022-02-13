package inkcaller

//InkState is based on Ink definition.
//We are ignoring many fields.
//This is only for reading. Writing variables into the states uses another way.
type InkState struct {
	Flows Flows `json:"flows"`
	//TurnIdx increases by 1 every call. Starts at 0.
	TurnIdx uint64 `json:"turnIdx"`
}

type Flows struct {
	DefaultFlow DefaultFlow `json:"DEFAULT_FLOW"`
}

type DefaultFlow struct {
	//OutputStream is the text
	OutputStream   []string        `json:"outputStream"`
	CurrentChoices []CurrentChoice `json:"currentChoices"`
}

type CurrentChoice struct {
	Text string `json:"text"`
	//Index is is the arg to pick this choice
	//20220212 This is always 0 in the exported JSON
	//https://github.com/y-lohse/inkjs/issues/932
	//Need fixup after decoding (use the position in the array)
	Index      uint16 `json:"index"`
	TargetPath string `json:"targetPath"`
}

func NewInkStateMock(turnIdx uint64, text []string, choices []CurrentChoice) *InkState {
	return &InkState{
		Flows:   Flows{DefaultFlow: DefaultFlow{CurrentChoices: choices, OutputStream: text}},
		TurnIdx: turnIdx,
	}
}
