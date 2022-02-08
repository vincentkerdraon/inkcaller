package inkcaller

import (
	"context"
)

type (
	ChoiceIndex int8
	KnotName    string

	Choice struct {
		Index ChoiceIndex
		Text  string
	}

	//Seed in ink can be 0<=seed<=100
	//(A modulo will be done)
	Seed int16

	InkCaller interface {
		Call(
			ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			seed *Seed,
			stateIn *StateEncoded,
			knotName *KnotName,
			choiceIndex *ChoiceIndex,
		) (*StateEncoded, error)
	}
)
