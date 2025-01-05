package biden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMarshalUnmarshal[T any](t *testing.T, value T, size int, marshaler MarshalFunc[T], unmarshaler UnmarshalFunc[T]) {
	encoded := Marshal(value, size, marshaler)
	decoded := Unmarshal(encoded, unmarshaler)

	assert.Equal(t, value, decoded)

	encoded2 := Marshal(decoded, size, marshaler)
	assert.Equal(t, encoded, encoded2)
}
