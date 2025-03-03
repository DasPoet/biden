package biden

type (
	MarshalFunc[T any]   func(pos int, buf []byte, value T) int
	UnmarshalFunc[T any] func(pos int, buf []byte) (T, int)
)

func Marshal[T any](value T, size int, marshaler MarshalFunc[T]) []byte {
	buf := make([]byte, size)
	marshaler(0, buf, value)
	return buf
}

func Unmarshal[T any](buf []byte, unmarshaler UnmarshalFunc[T]) T {
	value, _ := unmarshaler(0, buf)
	return value
}
