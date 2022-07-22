package inkcallerlib

import (
	"context"
	"math"
	"strconv"
)

type (
	//ChoiceIndex starts at 0
	ChoiceIndex int8

	KnotName string

	//SourcePath is a nearly ID for a choice (not unique!)
	SourcePath string

	Choice struct {
		Index ChoiceIndex
		//Text is always 1 line by ink design.
		Text       string
		SourcePath SourcePath
	}

	//Seed is for random. In ink it can be 0<=seed<=100
	//(A modulo will be done inside ink)
	Seed int8

	//TurnIndex is nearly an ID for the current call.
	//Starts at -1
	TurnIndex int16

	//Tag is a inkTag, using '#' syntax.
	//This is a GlobalTag or a LineTag.
	//Unfortunately, there is no ChoiceTag
	Tag string

	//Line is the output text (and tags).
	//Line.Tags is nil if not requested.
	Line struct {
		Text string
		Tags *[]Tag
	}

	InkCaller interface {
		Call(
			ctx context.Context,
			engineFilePath string,
			storyFilePath string,
			opts ...InkCallerOptionsFunc,
		) (*InkCallerOutput, error)
	}

	InkCallerOutput struct {
		//Lines is the text output. This is nil if not requested.
		Lines *[]Line
		//Choices are the player options. This is nil if not requested.
		Choices *[]Choice
		//StateOut is the ink json state. This is nil if not requested.
		StateOut *StateEncoded
		//GlobalTags are the story tags. This is nil if not requested.
		GlobalTags *[]Tag
		//TurnIndex increments every call. This is nil if not requested.
		TurnIndex *TurnIndex
	}
)

func NewSeedFromUint64(i uint64) Seed {
	return Seed(i % math.MaxInt8)
}

func (i TurnIndex) String() string {
	return strconv.Itoa(int(i))
}

func ParseTurnIndex(s string) (*TurnIndex, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	ti := TurnIndex(i)
	return &ti, err
}
