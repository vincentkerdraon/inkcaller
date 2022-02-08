package inktranslatormock

import (
	"context"

	"github.com/vincentkerdraon/inkcaller"
	"github.com/vincentkerdraon/inkcaller/inktranslator"
)

type (
	InkTranslatorMock struct {
		ResBeginStory      *inkcaller.InkState
		ResGoToKnot        *inkcaller.InkState
		ResContinueStory   *inkcaller.InkState
		ResGetResourceText []string
	}
)

func NewInkTranslatorMock() *InkTranslatorMock {
	return &InkTranslatorMock{}
}

var _ inktranslator.InkTranslator = (*InkTranslatorMock)(nil)

func (t *InkTranslatorMock) BeginStory(ctx context.Context, engineFilePath string, storyFilePath string, seed *inkcaller.Seed,
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	se := inkcaller.StateEncoded("")
	return t.ResBeginStory, &se, nil
}

func (t *InkTranslatorMock) GoToKnot(ctx context.Context, engineFilePath string, storyFilePath string, stateIn inkcaller.StateEncoded, gameModelV map[string]interface{}, knotName inkcaller.KnotName,
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	se := inkcaller.StateEncoded("")
	return t.ResGoToKnot, &se, nil
}

func (t *InkTranslatorMock) ContinueStory(ctx context.Context, engineFilePath string, storyFilePath string, choiceIndex inkcaller.ChoiceIndex, stateIn inkcaller.StateEncoded, gameModelV map[string]interface{},
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	se := inkcaller.StateEncoded("")
	return t.ResContinueStory, &se, nil
}

func (t *InkTranslatorMock) GetResourceText(ctx context.Context, engineFilePath string, storyFilePath string, knotName inkcaller.KnotName,
) ([]string, error) {
	return t.ResGetResourceText, nil
}
