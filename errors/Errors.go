package errors

type IndexOutOfRange struct{}

func (t IndexOutOfRange) Error() string {
    return "index out of range"
}
