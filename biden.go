package biden

const Int64Bytes = 8

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

func MarshalInt64(pos int, buf []byte, n int64) int {
	window := buf[pos : pos+Int64Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[7]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)
	window[4] = byte(n >> 32)
	window[5] = byte(n >> 40)
	window[6] = byte(n >> 48)
	window[7] = byte(n >> 56)

	return pos + Int64Bytes
}

func UnmarshalInt64(pos int, buf []byte) (int64, int) {
	window := buf[pos : pos+Int64Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[7]

	hi := uint64(window[0]) | uint64(window[1])<<8 | uint64(window[2])<<16 | uint64(window[3])<<24
	lo := uint64(window[4])<<32 | uint64(window[5])<<40 | uint64(window[6])<<48 | uint64(window[7])<<56

	return int64(hi | lo), pos + Int64Bytes
}
