package inkcallerlib

type (
	InkError struct{ Err error }
)

func (e InkError) Unwrap() error {
	return e.Err
}

func (e InkError) Error() string {
	return e.Err.Error()
}
