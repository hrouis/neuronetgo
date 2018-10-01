package math

type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	ErrShape               = Error{"matrix: dimension mismatch"}
)