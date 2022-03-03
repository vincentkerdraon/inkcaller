package inkcallerlib

// For now, keeping all that open. But with the InkCallerOptionsFunc syntax,
//everything can be made non-exported and using With.. functions to protect the usage.

type (
	InkCallerOptionsInput struct {
		Seed        *Seed
		StateIn     *StateEncoded
		KnotName    *KnotName
		ChoiceIndex *ChoiceIndex
	}
	InkCallerOptionsOutput struct {
		//all those options are not really necessary. The benchmark is not showing significant diff.
		//but more options will come.
		//I find the problem is mostly on the output object format (pointer are not fun to use).

		Choices    bool
		Lines      bool
		LineTags   bool
		StateOut   bool
		GlobalTags bool
		TurnIndex  bool
	}

	InkCallerOptions struct {
		Input  InkCallerOptionsInput
		Output InkCallerOptionsOutput
	}

	InkCallerOptionsFunc func(*InkCallerOptions)
)

func ReadOptions(opts []InkCallerOptionsFunc) InkCallerOptions {
	//At the begining, I was going to set Choices+Lines+StateOut true by default.
	//But I see this assumption is false.
	options := InkCallerOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

func WithInputSeed(seed Seed) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Input.Seed = &seed
	}
}

func WithInputStateIn(stateIn StateEncoded) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Input.StateIn = &stateIn
	}
}

func WithInputKnotName(knotName KnotName) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Input.KnotName = &knotName
	}
}

func WithInputChoiceIndex(choiceIndex ChoiceIndex) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Input.ChoiceIndex = &choiceIndex
	}
}

func WithOutputChoices(b bool) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Output.Choices = b
	}
}

func WithOutputLines(b bool) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Output.Lines = b
	}
}

func WithOutputLineTags(b bool) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Output.LineTags = b
	}
}

func WithOutputStateOut(b bool) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Output.StateOut = b
	}
}

func WithOutputGlobalTags(b bool) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Output.GlobalTags = b
	}
}

func WithOutputTurnIndex(b bool) InkCallerOptionsFunc {
	return func(ico *InkCallerOptions) {
		ico.Output.TurnIndex = b
	}
}
