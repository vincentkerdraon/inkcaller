package inkcallerv8

import "fmt"

type (
	InkV8Error struct {
		Source string
		Err    error
	}
)

func (e InkV8Error) Unwrap() error {
	return e.Err
}

func (e InkV8Error) Error() string {
	return fmt.Sprintf("V8 Error, %+v\n%s", e.Err.Error(), e.Source)
}
