# inkcaller

See also the godoc https://pkg.go.dev/github.com/vincentkerdraon/inkcaller

## This library

`InkCaller` is an API to call Ink. Each new call is independent and can be executing concurrently.\
A call will force the ink state, (optional) set the seed, (optional) set the knot, (optional) answer a previous choice.\
A call returns the current ink state (to inject into the next call). This state can be decoded to get the text, the choices.\
This does not allow all the features of ink (e.g. no callback / external function / multiple flows).

Depending on your design, Ink will need to interact with the game model. Set needed data inside ink state in the input. Use formatted text and a parser for the actions.

`InkTranslator` is a helper to decode the json ink state.

## Example

### Example using directly the inkcallerv8

```golang
	caller := inkcallerv8.NewInkCallerV8()

	//first call, we don't have an ink state to provide yet.
	//- set seed
	//- get ink JSON state + choices + no text
	out, err := caller.Call(ctx, engineFilePath, storyFilePath,
		//use WithInput... provide parameters
		inkcallerlib.WithInputSeed(2),
		//use WithOutput... pick what you want in the output
		inkcallerlib.WithOutputLines(false),
		inkcallerlib.WithOutputStateOut(true),
		inkcallerlib.WithOutputChoices(true))
	if err != nil {
		panic(err)
	}
	fmt.Println((*out.Choices)[0].Text)

	//second call, always use StateOut from the previous call.
	//- go to knot + make a choice
	//- get text + no choice
	out, err = caller.Call(ctx, engineFilePath, storyFilePath,
		inkcallerlib.WithInputStateIn(*out.StateOut),
		inkcallerlib.WithInputKnotName("Hub"),
		inkcallerlib.WithInputChoiceIndex(0),
		inkcallerlib.WithOutputChoices(false),
		inkcallerlib.WithOutputLines(true))

	if err != nil {
		panic(err)
	}
	if out.Choices != nil {
		panic("Maybe there are choices available, but out.Choices is nil because not requested")
	}
	fmt.Println((*out.Lines)[0].Text)
```

### Example using the translator, decoding the state

The `translator` package guides the possible operations and decodes the ink state into a go struct. This is using inkcallerv8 as a dependency (or any mock for easy integration tests).

```golang
	var inkTranslator InkTranslator
	ctx := context.Background()

	//First call
	out, err := inkTranslator.Begin(ctx, "ink.js", "story.json", inkcallerlib.Seed(0))
	if err != nil {
		panic(err)
	}

	//Using StateOut from previous call
	//A variable "Level" is declared in the ink file, overriding its value.
	gameData := map[string]interface{}{"Level": "5"}
	out, err = inkTranslator.GoToKnot(ctx, "ink.js", "story.json", *out.StateOut, gameData, inkcallerlib.KnotName("Hub"))
	if err != nil {
		panic(err)
	}

	//From the previous state, making a choice.
	out, err = inkTranslator.Decide(ctx, "ink.js", "story.json", *out.StateOut, gameData, inkcallerlib.ChoiceIndex(1))
	if err != nil {
		panic(err)
	}
```

## Advised flow for writing a story and using this lib

Write on the editor inky. (Create yourself a debug knot to test everything as you write. See `INK_DEBUG` in the example.)\
When you done writing, either use inky to `export as JSON`, or use the command line `inklecate -j -o story_demo.json story_demo.ink`.\
This lib can read directly your output .json file.\
You probably want to do this sequence `BeginStory() -> GoToKnot() -> ContinueStory() -> ContinueStory() -> ... -> GoToKnot() -> ContinueStory() ...`\
Or to get story-independent strings `GetResourceText()` (externalize all strings for easy translation and maintenance, example: menu, credits ...)

## What is ink?

"Ink" is a narrative scripting language for games. https://www.inklestudios.com/ink/

This is redacted in a `story.ink` source file. Then it can be exported for web: `story.js` + `ink.js`.\
Quick Start Ink: https://github.com/inkle/ink/blob/master/Documentation/WritingWithInk.md

Inkle projects and related:
- https://github.com/inkle/ink (the original engine in C#)
- https://github.com/y-lohse/inkjs (the JS portage, used in this project)
- https://github.com/inkle/inky (the editor and export js)
- https://github.com/inkle/ink-unity-integration (alternative to this project, direct integration in unity)

## License

MIT
