package inkcallerv8_test

import (
	"context"
	"path/filepath"

	"github.com/vincentkerdraon/inkcaller"
	"github.com/vincentkerdraon/inkcaller/inkcallerv8"
)

func Example() {
	ctx := context.Background()
	storyFilePath, _ := filepath.Abs("../assets_demo/story/story_demo.json")
	engineFilePath, _ := filepath.Abs("../assets_demo/ink_engine/1.0/ink.js")

	caller := inkcallerv8.NewInkCallerV8()

	//first call, we don't have an ink state to provide yet.
	//we also want to set the seed.
	//this is going to return the top part (introduction) in the ink file.

	seed := inkcaller.Seed(2)
	stateEncoded, err := caller.Call(ctx, engineFilePath, storyFilePath, &seed, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	//second call, use stateEncoded from the previous call.
	//go to knot
	//make a choice
	choice := inkcaller.ChoiceIndex(0)
	knotName := inkcaller.KnotName("Hub")
	stateEncoded, err = caller.Call(ctx, engineFilePath, storyFilePath, nil, stateEncoded, &knotName, &choice)
	if err != nil {
		panic(err)
	}

	_ = stateEncoded
}
