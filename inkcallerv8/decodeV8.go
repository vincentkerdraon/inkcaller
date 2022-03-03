package inkcallerv8

import (
	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
	"rogchap.com/v8go"
)

//All this decoding from v8 objects is doing a lot of allocation.
//This is also error prone code, hard to debug.
//Idea would be to generate JSON on the js side, get as a string, and then do the work on go side.
//20220301 perf using that is not too bad, 13% slower when decoding everything compare to no decode.
// - When no output. 3331378 ns/op. 36 allocs/op
// - When every outputs. 3758420 ns/op. 354 allocs/op

func (c *impl) v8DecodeArray(val *v8go.Value, rowFunc func(int32, *v8go.Value) error) error {
	//I have no idea how to use arrays.
	//Instead, consider the array an object and do it manually.

	arr := val.Object()
	len, err := arr.Get("length")
	if err != nil {
		return err
	}
	_len := len.Int32()
	for i := int32(0); i < _len; i++ {
		elem, err := arr.GetIdx(uint32(i))
		if err != nil {
			return err
		}
		if err := rowFunc(i, elem); err != nil {
			return err
		}
	}
	return nil
}

func (c *impl) v8DecodeChoices(valObj *v8go.Object) ([]inkcallerlib.Choice, error) {
	arrChoices, err := valObj.Get("Choices")
	if err != nil {
		return nil, err
	}
	var choices []inkcallerlib.Choice
	err = c.v8DecodeArray(arrChoices, func(_ int32, valChoice *v8go.Value) error {
		objChoice, err := valChoice.AsObject()
		if err != nil {
			return err
		}
		objChoiceText, err := objChoice.Get("Text")
		if err != nil {
			return err
		}
		objChoiceIndex, err := objChoice.Get("Index")
		if err != nil {
			return err
		}
		objChoiceSourcePath, err := objChoice.Get("SourcePath")
		if err != nil {
			return err
		}

		choices = append(choices, inkcallerlib.Choice{
			SourcePath: inkcallerlib.SourcePath(objChoiceSourcePath.String()),
			Index:      inkcallerlib.ChoiceIndex(objChoiceIndex.Int32()),
			Text:       objChoiceText.String(),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return choices, nil
}

func (c *impl) v8DecodeLines(options inkcallerlib.InkCallerOptions, valObj *v8go.Object) ([]inkcallerlib.Line, error) {
	arrLines, err := valObj.Get("Lines")
	if err != nil {
		return nil, err
	}
	var lines []inkcallerlib.Line
	err = c.v8DecodeArray(arrLines, func(_ int32, valLine *v8go.Value) error {
		objLine, err := valLine.AsObject()
		if err != nil {
			return err
		}
		var tags *[]inkcallerlib.Tag
		if options.Output.LineTags {
			valLineTags, err := objLine.Get("Tags")
			if err != nil {
				return err
			}
			var res []inkcallerlib.Tag
			err = c.v8DecodeArray(valLineTags, func(_ int32, valLine *v8go.Value) error {
				res = append(res, inkcallerlib.Tag(valLine.String()))
				return nil
			})
			if err != nil {
				return err
			}
			tags = &res
		}

		valLineText, err := objLine.Get("Text")
		if err != nil {
			return err
		}
		lines = append(lines, inkcallerlib.Line{
			Text: valLineText.String(),
			Tags: tags,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func (c *impl) v8DecodeTags(valObj *v8go.Object) ([]inkcallerlib.Tag, error) {
	arrLines, err := valObj.Get("GlobalTags")
	if err != nil {
		return nil, err
	}
	var tags []inkcallerlib.Tag
	err = c.v8DecodeArray(arrLines, func(_ int32, valLine *v8go.Value) error {
		tags = append(tags, inkcallerlib.Tag(valLine.String()))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (c *impl) decodeV8Res(options inkcallerlib.InkCallerOptions, val *v8go.Value) (*inkcallerlib.InkCallerOutput, error) {
	//assume js returns exactly the structure InkCallerOutput (field absent if not requested in options)

	res := &inkcallerlib.InkCallerOutput{}
	valObj := val.Object()

	if options.Output.Lines {
		lines, err := c.v8DecodeLines(options, valObj)
		if err != nil {
			return nil, err
		}
		res.Lines = &lines
	}

	if options.Output.Choices {
		choices, err := c.v8DecodeChoices(valObj)
		if err != nil {
			return nil, err
		}
		res.Choices = &choices
	}

	if options.Output.GlobalTags {
		tags, err := c.v8DecodeTags(valObj)
		if err != nil {
			return nil, err
		}
		res.GlobalTags = &tags
	}

	if options.Output.StateOut {
		valStateOut, err := valObj.Get("State")
		if err != nil {
			return nil, err
		}
		stateOut := inkcallerlib.StateEncoded(valStateOut.String())
		res.StateOut = &stateOut
	}

	if options.Output.TurnIndex {
		valTurnIndex, err := valObj.Get("TurnIndex")
		if err != nil {
			return nil, err
		}
		turnIndex := inkcallerlib.TurnIndex(valTurnIndex.Int32())
		res.TurnIndex = &turnIndex
	}

	return res, nil
}
