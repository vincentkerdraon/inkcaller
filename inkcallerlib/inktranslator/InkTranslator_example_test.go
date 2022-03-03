package inktranslator

import (
	"context"

	inkcallerlib "github.com/vincentkerdraon/inkcaller/inkcallerlib"
)

func Example() {
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
}
