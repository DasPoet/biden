package biden

func SliceBytes[T any](slice []T, itemSize int) int {
	return IntBytes + len(slice)*itemSize
}

func StringSliceBytes(strings []string) int {
	var contentBytes int
	for _, s := range strings {
		contentBytes += StringBytes(s)
	}
	return IntBytes + contentBytes
}

func MarshalSlice[T any](pos int, buf []byte, slice []T, marshaler MarshalFunc[T]) int {
	pos = MarshalInt(pos, buf, len(slice))
	for _, item := range slice {
		pos = marshaler(pos, buf, item)
	}
	return pos
}

func UnmarshalSlice[T any](pos int, buf []byte, unmarshaler UnmarshalFunc[T]) ([]T, int) {
	var (
		size int
		item T
	)

	size, pos = UnmarshalInt(pos, buf)
	items := make([]T, size)

	for i := range size {
		item, pos = unmarshaler(pos, buf)
		items[i] = item
	}
	return items, pos
}
