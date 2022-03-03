package inkcallerv8_test

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
	"github.com/vincentkerdraon/inkcaller/inkcallerv8"
)

func Example() {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/inkjs_engine/2.0.0/ink.js")

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
}
