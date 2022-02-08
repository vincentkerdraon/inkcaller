package inkcallerv8_test

import (
	"context"
	"fmt"

	"github.com/vincentkerdraon/inkcaller"
	"github.com/vincentkerdraon/inkcaller/inkcallerv8"
	"github.com/vincentkerdraon/inkcaller/inktranslator"
)

func Example2() {
	ctx := context.Background()
	caller := inkcallerv8.NewInkCallerV8()
	//translator guides the possible operations and decodes the output
	translator := inktranslator.NewInkTranslator(caller)

	seed := inkcaller.Seed(2)
	inkState, inkStateEncoded, err := translator.BeginStory(ctx, "ink.js", "story.json", &seed)
	if err != nil {
		panic(err)
	}
	fmt.Println(inkState.OutputStream)

	knotName := inkcaller.KnotName("Hub")
	inkState, inkStateEncoded, err = translator.GoToKnot(ctx, "ink.js", "story.json", *inkStateEncoded, map[string]interface{}{}, knotName)
	if err != nil {
		panic(err)
	}
	fmt.Println(inkState.CurrentChoices[0].Text)

	_ = inkStateEncoded
}
