package biden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Item struct {
	Name   string
	Height uint8
}

func ItemBytes(i Item) int {
	return StringBytes(i.Name) + Uint8Bytes
}

func TestDynamicMapBytes(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var (
			m     = map[string]Item{}
			bytes = DynamicMapBytes(m, StringBytes, ItemBytes)
		)
		assert.Equal(t, IntBytes, bytes)
	})

	t.Run("full map", func(t *testing.T) {
		var (
			m = map[string]Item{
				"bob": {
					Name:   "bryan",
					Height: 185,
				},
			}
			bytes = DynamicMapBytes(m, StringBytes, ItemBytes)
		)
		assert.Equal(t, IntBytes+IntBytes+len("bob")+IntBytes+len("bryan")+Uint8Bytes, bytes)
	})
}

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
