package inkcallerlib

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StateEncoded string

func (sEncoded *StateEncoded) IncludeGameData(gameModelMV map[string]interface{}) error {
	//this is a bit tricky.
	//state.VariablesState contains multiple things.
	// - (keep untouched the ink internal variables, and the rest of the state)
	// - remove everything related to the game model
	// - insert the provided game model

	//We could use a variable prefix, that we could always delete.
	//Instead, assume gameModelV is always having all the keys.
	//even to empty, the key should be present.

	bEncoded := []byte(*sEncoded)

	state := make(map[string]interface{})
	if err := json.Unmarshal(bEncoded, &state); err != nil {
		return err
	}

	var inkStateOnlyVariable struct {
		VariablesState json.RawMessage `json:"variablesState"`
	}
	if err := json.Unmarshal(bEncoded, &inkStateOnlyVariable); err != nil {
		return err
	}

	inkVariables := make(map[string]interface{})
	if err := json.Unmarshal(inkStateOnlyVariable.VariablesState, &inkVariables); err != nil {
		return err
	}

	for k, v := range gameModelMV {
		inkVariables[k] = v
	}

	b, err := json.Marshal(inkVariables)
	if err != nil {
		return err
	}

	state["variablesState"] = json.RawMessage(b)
	bState, err := json.Marshal(state)
	if err != nil {
		return err
	}

	*sEncoded = StateEncoded(bState)
	return nil
}

//ConvertVtoMV is a helper to transform any structure in a suitable game data interface
func (StateEncoded) ConvertVtoMV(gameModelV interface{}) (map[string]interface{}, error) {
	//(using reflexion, so json conventions won't matter)

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
