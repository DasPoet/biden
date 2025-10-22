package biden

import "testing"

func TestMarshalUnmarshalMap(t *testing.T) {
	testMarshalUnmarshalMap(t, map[rune]bool{}, RuneBytes, BoolBytes, MarshalInt32, MarshalBool, UnmarshalInt32, UnmarshalBool)

	testMarshalUnmarshalMap(t, map[rune]bool{'a': true, 'b': false}, RuneBytes, BoolBytes, MarshalInt32, MarshalBool, UnmarshalInt32, UnmarshalBool)
	testMarshalUnmarshalMap(t, map[rune]bool{'a': false, 'b': true}, RuneBytes, BoolBytes, MarshalInt32, MarshalBool, UnmarshalInt32, UnmarshalBool)
}

func testMarshalUnmarshalMap[K comparable, V any](t *testing.T, m map[K]V, keySize, valueSize int, keyMarshaler MarshalFunc[K], valueMarshaler MarshalFunc[V], keyUnmarshaler UnmarshalFunc[K], valueUnmarshaler UnmarshalFunc[V]) {
	var (
		size = MapBytes(m, keySize, valueSize)

		marshaler   = NewMapMarshaler(keyMarshaler, valueMarshaler)
		unmarshaler = NewMapUnmarshaler(keyUnmarshaler, valueUnmarshaler)
	)

	testMarshalUnmarshal(t, m, size, marshaler, unmarshaler)
}
