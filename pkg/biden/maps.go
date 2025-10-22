package biden

func MapBytes[K comparable, V any](m map[K]V, keySize, valueSize int) int {
	return IntBytes + len(m)*(keySize+valueSize)
}

func DynamicMapBytes[K comparable, V any](m map[K]V, keySizer SizerFunc[K], valueSizer SizerFunc[V]) int {
	size := IntBytes
	for key, value := range m {
		size += keySizer(key)
		size += valueSizer(value)
	}
	return size
}

func MarshalMap[K comparable, V any](pos int, buf []byte, m map[K]V, marshalKey MarshalFunc[K], marshalValue MarshalFunc[V]) int {
	pos = MarshalInt(pos, buf, len(m))
	for key, value := range m {
		pos = marshalKey(pos, buf, key)
		pos = marshalValue(pos, buf, value)
	}
	return pos
}

func UnmarshalMap[K comparable, V any](pos int, buf []byte, unmarshalKey UnmarshalFunc[K], unmarshalValue UnmarshalFunc[V]) (map[K]V, int) {
	var (
		size  int
		key   K
		value V
	)

	size, pos = UnmarshalInt(pos, buf)
	result := make(map[K]V, size)

	for range size {
		key, pos = unmarshalKey(pos, buf)
		value, pos = unmarshalValue(pos, buf)

		result[key] = value
	}
	return result, pos
}
