package inktranslator

import (
	"context"

	"github.com/vincentkerdraon/inkcaller"
)

type (
	InkTranslator interface {
		BeginStory(ctx context.Context, engineFilePath string, storyFilePath string, seed *inkcaller.Seed,
		) (*inkcaller.InkState, *inkcaller.StateEncoded, error)

		GoToKnot(ctx context.Context, engineFilePath string, storyFilePath string, stateIn inkcaller.StateEncoded, gameModelV map[string]interface{}, knotName inkcaller.KnotName,
		) (*inkcaller.InkState, *inkcaller.StateEncoded, error)

		ContinueStory(ctx context.Context, engineFilePath string, storyFilePath string, choiceIndex inkcaller.ChoiceIndex, stateIn inkcaller.StateEncoded, gameModelV map[string]interface{},
		) (*inkcaller.InkState, *inkcaller.StateEncoded, error)

		GetResourceText(ctx context.Context, engineFilePath string, storyFilePath string, knotName inkcaller.KnotName,
		) ([]string, error)
	}

	impl struct {
		inkCaller inkcaller.InkCaller
	}
)

func NewInkTranslator(inkCaller inkcaller.InkCaller) *impl {
	return &impl{
		inkCaller: inkCaller,
	}
}

var _ InkTranslator = (*impl)(nil)

func (t *impl) BeginStory(ctx context.Context, engineFilePath string, storyFilePath string, seed *inkcaller.Seed,
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	return t.callAndDecode(ctx, engineFilePath, storyFilePath, seed, nil, nil, nil, nil)
}

func (t *impl) GoToKnot(ctx context.Context, engineFilePath string, storyFilePath string, stateIn inkcaller.StateEncoded, gameModelV map[string]interface{}, knotName inkcaller.KnotName,
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	return t.callAndDecode(ctx, engineFilePath, storyFilePath, nil, &stateIn, gameModelV, &knotName, nil)
}

func (t *impl) ContinueStory(ctx context.Context, engineFilePath string, storyFilePath string, choiceIndex inkcaller.ChoiceIndex, stateIn inkcaller.StateEncoded, gameModelV map[string]interface{},
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	return t.callAndDecode(ctx, engineFilePath, storyFilePath, nil, &stateIn, gameModelV, nil, &choiceIndex)
}

func (t *impl) GetResourceText(ctx context.Context, engineFilePath string, storyFilePath string, knotName inkcaller.KnotName,
) ([]string, error) {
	state, _, err := t.callAndDecode(ctx, engineFilePath, storyFilePath, nil, nil, nil, &knotName, nil)
	if err != nil {
		return nil, err
	}
	return state.OutputStream, nil
}

func (t *impl) callAndDecode(ctx context.Context, engineFilePath string, storyFilePath string, seed *inkcaller.Seed, stateIn *inkcaller.StateEncoded, gameModelV map[string]interface{}, knotName *inkcaller.KnotName, choiceIndex *inkcaller.ChoiceIndex,
) (*inkcaller.InkState, *inkcaller.StateEncoded, error) {
	if stateIn != nil && gameModelV != nil {
		var err error
		stateIn, err = stateIn.IncludeGameData(gameModelV)
		if err != nil {
			return nil, nil, inkcaller.InkError{Err: err}
		}
	}

	stateEncoded, err := t.inkCaller.Call(ctx, engineFilePath, storyFilePath, seed, stateIn, knotName, choiceIndex)
	if err != nil {
		return nil, nil, inkcaller.InkError{Err: err}
	}
	state, err := stateEncoded.DecodeInkState()
	if err != nil {
		return nil, nil, inkcaller.InkError{Err: err}
	}
	return state, stateEncoded, nil
}
