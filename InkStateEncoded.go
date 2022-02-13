package inkcaller

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StateEncoded string

func (sEncoded StateEncoded) DecodeInkState() (*InkState, error) {
	state := &InkState{}
	err := json.Unmarshal([]byte(sEncoded), state)
	if err != nil {
		return nil, err
	}

	//20220212 Index is always 0 in the exported JSON
	//https://github.com/y-lohse/inkjs/issues/932
	//Need fixup after decoding (use the position in the array)
	for idx := range state.Flows.DefaultFlow.CurrentChoices {
		state.Flows.DefaultFlow.CurrentChoices[idx].Index = uint16(idx)
	}

	return state, nil
}

func (sEncoded StateEncoded) IncludeGameData(gameModelMV map[string]interface{}) (*StateEncoded, error) {
	//this is a bit tricky.
	//state.VariablesState contains multiple things.
	// - (keep untouched the ink internal variables, and the rest of the state)
	// - remove everything related to the game model
	// - insert the provided game model

	//We could use a variable prefix, that we could always delete.
	//Instead, assume gameModelV is always having all the keys.
	//even if empty, the key should be present.
	//(using reflexion, so json conventions won't matter)

	bEncoded := []byte(sEncoded)

	state := make(map[string]interface{})
	if err := json.Unmarshal(bEncoded, &state); err != nil {
		return nil, err
	}

	var inkStateOnlyVariable struct {
		VariablesState json.RawMessage `json:"variablesState"`
	}
	if err := json.Unmarshal(bEncoded, &inkStateOnlyVariable); err != nil {
		return nil, err
	}

	inkVariables := make(map[string]interface{})
	if err := json.Unmarshal(inkStateOnlyVariable.VariablesState, &inkVariables); err != nil {
		return nil, err
	}

	for k, v := range gameModelMV {
		inkVariables[k] = v
	}

	b, err := json.Marshal(inkVariables)
	if err != nil {
		return nil, err
	}

	state["variablesState"] = json.RawMessage(b)

	bState, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	sState := StateEncoded(string(bState))
	return &sState, nil
}

func (StateEncoded) ConvertVtoMV(gameModelV interface{}) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	v := reflect.ValueOf(gameModelV)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			res[v.Type().Field(i).Name] = v.Field(i).Interface()
		} else {
			//blocking the possibility instead of having strange bugs
			return nil, fmt.Errorf("fail reflect on un-exported type")
		}
	}
	return res, nil
}
