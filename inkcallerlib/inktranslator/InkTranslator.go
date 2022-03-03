//package inktranslator is a guide to use InkCaller.
package inktranslator

import (
	"context"

	inkcallerlib "github.com/vincentkerdraon/inkcaller/inkcallerlib"
)

type (
	//InkTranslator is a guide to use InkCaller (for a specific usage).
	//Prefer using directly inkcallerlib.InkCaller for your usecase.
	InkTranslator interface {

		//Begin is to initialize the ink state.
		//Set the random seed, maybe check the story meta?
		Begin(
			ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			seed inkcallerlib.Seed,
		) (*inkcallerlib.InkCallerOutput, error)

		//GoToKnot will visit a knot. Maybe the "Hub" where all starts?
		//And return all possible starters from this point.
		GoToKnot(
			ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			stateIn inkcallerlib.StateEncoded,
			gameModelV map[string]interface{},
			knotName inkcallerlib.KnotName,
		) (*inkcallerlib.InkCallerOutput, error)

		//Decide will input the player choice and continue the story until the next choice (or the end).
		Decide(ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			stateIn inkcallerlib.StateEncoded,
			gameModelV map[string]interface{},
			choiceIndex inkcallerlib.ChoiceIndex,
		) (*inkcallerlib.InkCallerOutput, error)

		//ResourceDynamicText is to get a complex text (with variables) from ink. E.g. a translation that depends on the context.
		ResourceDynamicText(
			ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			stateIn inkcallerlib.StateEncoded,
			gameModelV map[string]interface{},
			knotName inkcallerlib.KnotName,
		) ([]inkcallerlib.Line, error)

		//ResourceStaticText is a text that never changes. E.g. menu entry ...
		ResourceStaticText(
			ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			knotName inkcallerlib.KnotName,
		) ([]inkcallerlib.Line, error)
	}

	impl struct {
		inkCaller inkcallerlib.InkCaller
	}
)

func NewInkTranslator(inkCaller inkcallerlib.InkCaller) *impl {
	return &impl{
		inkCaller: inkCaller,
	}
}

var _ InkTranslator = (*impl)(nil)

func (t *impl) Begin(
	ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	seed inkcallerlib.Seed,
) (*inkcallerlib.InkCallerOutput, error) {
	out, err := t.inkCaller.Call(ctx, engineFilePath, storyFilePath,
		inkcallerlib.WithInputSeed(seed),
		inkcallerlib.WithOutputGlobalTags(true),
		inkcallerlib.WithOutputStateOut(true),
	)
	if err != nil {
		return nil, inkcallerlib.InkError{Err: err}
	}
	return out, nil
}

func (t *impl) GoToKnot(
	ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	stateIn inkcallerlib.StateEncoded,
	gameModelV map[string]interface{},
	knotName inkcallerlib.KnotName,
) (*inkcallerlib.InkCallerOutput, error) {
	if err := stateIn.IncludeGameData(gameModelV); err != nil {
		return nil, err
	}
	out, err := t.inkCaller.Call(ctx, engineFilePath, storyFilePath,
		inkcallerlib.WithInputStateIn(stateIn),
		inkcallerlib.WithInputKnotName(knotName),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputLines(true),
	)
	if err != nil {
		return nil, inkcallerlib.InkError{Err: err}
	}
	return out, nil
}

func (t *impl) Decide(ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	stateIn inkcallerlib.StateEncoded,
	gameModelV map[string]interface{},
	choiceIndex inkcallerlib.ChoiceIndex,
) (*inkcallerlib.InkCallerOutput, error) {
	if err := stateIn.IncludeGameData(gameModelV); err != nil {
		return nil, err
	}
	out, err := t.inkCaller.Call(ctx, engineFilePath, storyFilePath,
		inkcallerlib.WithInputStateIn(stateIn),
		inkcallerlib.WithInputChoiceIndex(choiceIndex),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputLines(true),
		inkcallerlib.WithOutputLineTags(true),
	)
	if err != nil {
		return nil, inkcallerlib.InkError{Err: err}
	}
	return out, nil
}

func (t *impl) ResourceDynamicText(
	ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	stateIn inkcallerlib.StateEncoded,
	gameModelV map[string]interface{},
	knotName inkcallerlib.KnotName,
) ([]inkcallerlib.Line, error) {
	if err := stateIn.IncludeGameData(gameModelV); err != nil {
		return nil, err
	}
	out, err := t.inkCaller.Call(ctx, engineFilePath, storyFilePath,
		inkcallerlib.WithInputStateIn(stateIn),
		inkcallerlib.WithInputKnotName(knotName),
		inkcallerlib.WithOutputLines(true),
		inkcallerlib.WithOutputLineTags(true),
	)
	if err != nil {
		return nil, inkcallerlib.InkError{Err: err}
	}
	return *out.Lines, nil
}

func (t *impl) ResourceStaticText(
	ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	knotName inkcallerlib.KnotName,
) ([]inkcallerlib.Line, error) {
	out, err := t.inkCaller.Call(ctx, engineFilePath, storyFilePath,
		inkcallerlib.WithInputKnotName(knotName),
		inkcallerlib.WithOutputLines(true),
		inkcallerlib.WithOutputLineTags(true),
	)
	if err != nil {
		return nil, inkcallerlib.InkError{Err: err}
	}
	return *out.Lines, nil
}
