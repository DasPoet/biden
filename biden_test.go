package biden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshalInt64(t *testing.T) {
	testMarshalUnmarshal(t, 1, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 256, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 384, Int64Bytes, MarshalInt64, UnmarshalInt64)

	testMarshalUnmarshal(t, -1, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -256, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -384, Int64Bytes, MarshalInt64, UnmarshalInt64)
}

func testMarshalUnmarshal[T any](t *testing.T, value T, size int, marshaler MarshalFunc[T], unmarshaler UnmarshalFunc[T]) {
	encoded := Marshal(value, size, marshaler)
	decoded := Unmarshal(encoded, unmarshaler)

	assert.Equal(t, value, decoded)

	encoded2 := Marshal(decoded, size, marshaler)
	assert.Equal(t, encoded, encoded2)
}
