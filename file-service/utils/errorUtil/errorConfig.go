package errorUtil

type StringError struct {
	msg string
}

func (e StringError) Error() string {
	return e.msg
}
func NewStringError(msg string) error {
	return &StringError{msg: msg}
}
