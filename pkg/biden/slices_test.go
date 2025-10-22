package biden

import (
	"testing"

	"github.com/google/uuid"
)

func TestMarshalUnmarshalSlice(t *testing.T) {
	testMarshalUnmarshalSlice(t, []any{}, 1, nil, nil)

	testMarshalUnmarshalSlice(t, []bool{false, true}, BoolBytes, MarshalBool, UnmarshalBool)

	testMarshalUnmarshalSlice(t, []uint{1, 2, 3}, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshalSlice(t, []uint8{1, 2, 3}, Uint8Bytes, MarshalUint8, UnmarshalUint8)
	testMarshalUnmarshalSlice(t, []uint16{1, 2, 3}, Uint16Bytes, MarshalUint16, UnmarshalUint16)
	testMarshalUnmarshalSlice(t, []uint32{1, 2, 3}, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshalSlice(t, []uint64{1, 2, 3}, Uint64Bytes, MarshalUint64, UnmarshalUint64)

	testMarshalUnmarshalSlice(t, []int{1, -2, 3}, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshalSlice(t, []int8{1, -2, 3}, Int8Bytes, MarshalInt8, UnmarshalInt8)
	testMarshalUnmarshalSlice(t, []int16{1, -2, 3}, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshalSlice(t, []int32{1, -2, 3}, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshalSlice(t, []int64{1, -2, 3}, Int64Bytes, MarshalInt64, UnmarshalInt64)

	testMarshalUnmarshalSlice(t, []float32{-1, 1, -2.22, 3.99}, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshalSlice(t, []float64{-1, 1, -2.22, 3.99}, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
}

func TestMarshalUnmarshalStringSlice(t *testing.T) {
	var (
		data = []string{"", "hello", "world", ""}

		marshaler   = NewSliceMarshaler(MarshalString)
		unmarshaler = NewSliceUnmarshaler(UnmarshalString)
	)

	testMarshalUnmarshal(t, data, StringSliceBytes(data), marshaler, unmarshaler)
}

func TestMarshalUnmarshalUUIDSlice(t *testing.T) {
	var (
		data = []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}

		marshaler   = NewSliceMarshaler(MarshalUUID)
		unmarshaler = NewSliceUnmarshaler(UnmarshalUUID)
	)

	testMarshalUnmarshal(t, data, SliceBytes(data, UUIDBytes), marshaler, unmarshaler)
}

func testMarshalUnmarshalSlice[T any](t *testing.T, value []T, itemSize int, itemMarshaler MarshalFunc[T], itemUnmarshaler UnmarshalFunc[T]) {
	var (
		size = SliceBytes(value, itemSize)

		marshaler   = NewSliceMarshaler(itemMarshaler)
		unmarshaler = NewSliceUnmarshaler(itemUnmarshaler)
	)

	testMarshalUnmarshal(t, value, size, marshaler, unmarshaler)
}
