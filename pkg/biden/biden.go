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

func NewSliceMarshaler[T any](marshalItem MarshalFunc[T]) MarshalFunc[[]T] {
	return func(pos int, buf []byte, value []T) int {
		return MarshalSlice(pos, buf, value, marshalItem)
	}
}

func NewSliceUnmarshaler[T any](unmarshalItem UnmarshalFunc[T]) UnmarshalFunc[[]T] {
	return func(pos int, buf []byte) ([]T, int) {
		return UnmarshalSlice(pos, buf, unmarshalItem)
	}
}
