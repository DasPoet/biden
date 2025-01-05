package biden

import "testing"

func TestMarshalUnmarshalUint8(t *testing.T) {
	testMarshalUnmarshal(t, 0, Uint8Bytes, MarshalUint8, UnmarshalUint8)
	testMarshalUnmarshal(t, 1, Uint8Bytes, MarshalUint8, UnmarshalUint8)
	testMarshalUnmarshal(t, 10, Uint8Bytes, MarshalUint8, UnmarshalUint8)
	testMarshalUnmarshal(t, 255, Uint8Bytes, MarshalUint8, UnmarshalUint8)
}

func TestMarshalUnmarshalUint16(t *testing.T) {
	testMarshalUnmarshal(t, 0, Uint16Bytes, MarshalUint16, UnmarshalUint16)
	testMarshalUnmarshal(t, 1, Uint16Bytes, MarshalUint16, UnmarshalUint16)
	testMarshalUnmarshal(t, 10, Uint16Bytes, MarshalUint16, UnmarshalUint16)
	testMarshalUnmarshal(t, 255, Uint16Bytes, MarshalUint16, UnmarshalUint16)
	testMarshalUnmarshal(t, 256, Uint16Bytes, MarshalUint16, UnmarshalUint16)
	testMarshalUnmarshal(t, 1000, Uint16Bytes, MarshalUint16, UnmarshalUint16)
}

func TestMarshalUnmarshalUint32(t *testing.T) {
	testMarshalUnmarshal(t, 0, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshal(t, 1, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshal(t, 10, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshal(t, 255, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshal(t, 256, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshal(t, 1000, Uint32Bytes, MarshalUint32, UnmarshalUint32)
	testMarshalUnmarshal(t, 25_000, Uint32Bytes, MarshalUint32, UnmarshalUint32)
}

func TestMarshalUnmarshalUint64(t *testing.T) {
	testMarshalUnmarshal(t, 0, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 1, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 10, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 255, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 256, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 1000, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 25_000, Uint64Bytes, MarshalUint64, UnmarshalUint64)
	testMarshalUnmarshal(t, 100_000, Uint64Bytes, MarshalUint64, UnmarshalUint64)
}

func TestMarshalUnmarshalUint(t *testing.T) {
	testMarshalUnmarshal(t, 0, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 1, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 10, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 255, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 256, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 1000, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 25_000, UintBytes, MarshalUint, UnmarshalUint)
	testMarshalUnmarshal(t, 100_000, UintBytes, MarshalUint, UnmarshalUint)
}

func TestMarshalUnmarshalInt8(t *testing.T) {
	testMarshalUnmarshal(t, 0, Int8Bytes, MarshalInt8, UnmarshalInt8)
	testMarshalUnmarshal(t, 1, Int8Bytes, MarshalInt8, UnmarshalInt8)
	testMarshalUnmarshal(t, 10, Int8Bytes, MarshalInt8, UnmarshalInt8)
	testMarshalUnmarshal(t, 127, Int8Bytes, MarshalInt8, UnmarshalInt8)

	testMarshalUnmarshal(t, -1, Int8Bytes, MarshalInt8, UnmarshalInt8)
	testMarshalUnmarshal(t, -10, Int8Bytes, MarshalInt8, UnmarshalInt8)
	testMarshalUnmarshal(t, -128, Int8Bytes, MarshalInt8, UnmarshalInt8)
}

func TestMarshalUnmarshalInt16(t *testing.T) {
	testMarshalUnmarshal(t, 0, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, 1, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, 10, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, 255, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, 256, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, 1000, Int16Bytes, MarshalInt16, UnmarshalInt16)

	testMarshalUnmarshal(t, -1, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, -10, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, -255, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, -256, Int16Bytes, MarshalInt16, UnmarshalInt16)
	testMarshalUnmarshal(t, -1000, Int16Bytes, MarshalInt16, UnmarshalInt16)
}

func TestMarshalUnmarshalInt32(t *testing.T) {
	testMarshalUnmarshal(t, 0, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, 1, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, 10, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, 255, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, 256, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, 1000, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, 25_000, Int32Bytes, MarshalInt32, UnmarshalInt32)

	testMarshalUnmarshal(t, -1, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, -10, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, -255, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, -256, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, -1000, Int32Bytes, MarshalInt32, UnmarshalInt32)
	testMarshalUnmarshal(t, -25_000, Int32Bytes, MarshalInt32, UnmarshalInt32)
}

func TestMarshalUnmarshalInt64(t *testing.T) {
	testMarshalUnmarshal(t, 0, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 1, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 10, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 255, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 256, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 1000, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 25_000, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, 100_000, Int64Bytes, MarshalInt64, UnmarshalInt64)

	testMarshalUnmarshal(t, -1, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -10, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -255, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -256, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -1000, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -25_000, Int64Bytes, MarshalInt64, UnmarshalInt64)
	testMarshalUnmarshal(t, -100_000, Int64Bytes, MarshalInt64, UnmarshalInt64)
}

func TestMarshalUnmarshalInt(t *testing.T) {
	testMarshalUnmarshal(t, 0, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 1, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 10, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 255, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 256, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 1000, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 25_000, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, 100_000, IntBytes, MarshalInt, UnmarshalInt)

	testMarshalUnmarshal(t, -1, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, -10, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, -255, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, -256, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, -1000, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, -25_000, IntBytes, MarshalInt, UnmarshalInt)
	testMarshalUnmarshal(t, -100_000, IntBytes, MarshalInt, UnmarshalInt)
}

func TestMarshalUnmarshalFloat32(t *testing.T) {
	testMarshalUnmarshal(t, 0, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 1, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 10, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 255, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 256, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 1000, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 25_000, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 100_000, Float32Bytes, MarshalFloat32, UnmarshalFloat32)

	testMarshalUnmarshal(t, -1, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -10, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -255, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -256, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -1000, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -25_000, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -100_000, Float32Bytes, MarshalFloat32, UnmarshalFloat32)

	testMarshalUnmarshal(t, 1.12, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -10.04, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 255.09, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -256.99, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 1000.56, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, -25_000.33, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
	testMarshalUnmarshal(t, 100_000.19, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
}

func TestMarshalUnmarshalFloat64(t *testing.T) {
	testMarshalUnmarshal(t, 0, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 1, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 10, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 255, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 256, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 1000, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 25_000, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 100_000, Float64Bytes, MarshalFloat64, UnmarshalFloat64)

	testMarshalUnmarshal(t, -1, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -10, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -255, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -256, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -1000, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -25_000, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -100_000, Float64Bytes, MarshalFloat64, UnmarshalFloat64)

	testMarshalUnmarshal(t, 1.12, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -10.04, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 255.09, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -256.99, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 1000.56, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, -25_000.33, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
	testMarshalUnmarshal(t, 100_000.19, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
}

func TestMarshalUnmarshalString(t *testing.T) {
	testMarshalUnmarshal(t, "", StringBytes(""), MarshalString, UnmarshalString)
	testMarshalUnmarshal(t, "hello world", StringBytes("hello world"), MarshalString, UnmarshalString)
}
