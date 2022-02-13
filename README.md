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
ctx := context.Background()
caller := inkcallerv8.NewInkCallerV8()

//first call, we don't have any ink state to provide yet.
//we also want to set the seed.
//this is going to return the top part (introduction) from the ink file.

seed := inkcaller.Seed(2)
stateEncoded, err := caller.Call(ctx, "ink.js", "story.json", &seed, nil, nil, nil)
if err != nil {panic(err)}

//second call, use stateEncoded from the previous call.
//go to knot
//make a choice
choice := inkcaller.ChoiceIndex(0)
knotName := inkcaller.KnotName("Hub")
stateEncoded, err = caller.Call(ctx, "ink.js", "story.json", nil, stateEncoded, &knotName, &choice)
if err != nil {panic(err)}
```

### Example using the translator, decoding the state

The `translator` package guides the possible operations and decodes the ink state into a go struct. This is using inkcallerv8 as a dependency (or any mock for easy integration tests).

```golang
ctx := context.Background()
caller := inkcallerv8.NewInkCallerV8()
translator := inktranslator.NewInkTranslator(caller)

seed := inkcaller.Seed(2)
inkState, inkStateEncoded, err := translator.BeginStory(ctx, "ink.js", "story.json", &seed)
if err != nil {panic(err)}
fmt.Println(inkState.OutputStream)

knotName := inkcaller.KnotName("Hub")
//A variable "Level" is declared in the ink file, overriding its value. 
gameData:= map[string]interface{}{"Level":"5"}
inkState, inkStateEncoded, err = translator.GoToKnot(ctx, "ink.js", "story.json", *inkStateEncoded, gameData, knotName)
if err != nil {panic(err)}
fmt.Println(inkState.CurrentChoices[0].Text)
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
