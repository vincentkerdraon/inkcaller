package inkcallermock

import (
	"context"

	"github.com/vincentkerdraon/inkcaller/inkcallerlib"
)

//InkCallerMock is a helper for tests. Return what you wish.
type InkCallerMock struct {
	ResFunc func(inkcallerlib.InkCallerOptions) (*inkcallerlib.InkCallerOutput, error)
}

var _ inkcallerlib.InkCaller = (*InkCallerMock)(nil)

func (icm InkCallerMock) Call(
	ctx context.Context,
	engineFilePath string,
	storyFilePath string,
	opts ...inkcallerlib.InkCallerOptionsFunc,
) (*inkcallerlib.InkCallerOutput, error) {
	options := inkcallerlib.ReadOptions(opts)
	return icm.ResFunc(options)
}
