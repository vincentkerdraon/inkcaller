package inkcallerv8

import (
	"fmt"

	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
)

func (*impl) prepareJS(options inkcallerlib.InkCallerOptions) string {
	execJs := `
var story = new inkjs.Story(storyContent); 
res = {};
`
	if options.Input.Seed != nil {
		seed := *options.Input.Seed
		//ink limitation
		seed = seed % 100
		execJs += fmt.Sprintf(`story.state.storySeed=%d;
`, seed)
	}
	if options.Input.StateIn != nil {
		execJs += fmt.Sprintf(`story.state.LoadJson(%q);
`, *options.Input.StateIn)
	}
	if options.Input.KnotName != nil {
		execJs += fmt.Sprintf(`story.ChoosePathString(%q);
`, *options.Input.KnotName)
	}
	if options.Input.ChoiceIndex != nil {
		execJs += fmt.Sprintf(`story.ChooseChoiceIndex(%d);
`, *options.Input.ChoiceIndex)
	}

	//Must be executed to get choices
	execJs += `
lines = [];
while (story.canContinue) {
	line={};
	line.Text=story.Continue();`
	if options.Output.Lines {
		if options.Output.LineTags {
			execJs += `
	line.Tags=story.currentTags;`
		}
		execJs += `
	lines.push(line)`
	}
	execJs += `
}`
	if options.Output.Lines {
		execJs += `
res.Lines=lines;
`
	}

	if options.Output.Choices {
		execJs += `
choices = [];
story.currentChoices.forEach(function (choice) {
	choices.push({
		SourcePath: choice.sourcePath,
		Index: choice.index,
		Text: choice.text
	});
});
res.Choices=choices;
`
	}
	if options.Output.StateOut {
		execJs += `
res.State = story.state.toJson();`
	}
	if options.Output.TurnIndex {
		execJs += `
res.TurnIndex = story.state.currentTurnIndex;`
	}
	if options.Output.GlobalTags {
		execJs += `
res.GlobalTags = story.globalTags;`
	}

	execJs += `
res`
	return execJs
}
