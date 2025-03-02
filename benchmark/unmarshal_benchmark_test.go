package benchmark

import (
	biden "github.com/daspoet/biden/v1"
	"testing"
)

var (
	decodedUints   []uint
	decodedUint8s  []uint8
	decodedUint16s []uint16
	decodedUint32s []uint32
	decodedUint64s []uint64

	decodedInts   []int
	decodedInt8s  []int8
	decodedInt16s []int16
	decodedInt32s []int32
	decodedInt64s []int64

	decodedFloat32s []float32
	decodedFloat64s []float64

	decodedStrings []string
)

func BenchmarkUnmarshalUintSlice(b *testing.B) {
	nums := makeUints()
	decodedUints = benchmarkUnmarshalSlice(b, nums, biden.UintBytes, biden.MarshalUint, biden.UnmarshalUint)
}

func BenchmarkUnmarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	decodedUint8s = benchmarkUnmarshalSlice(b, nums, biden.Uint8Bytes, biden.MarshalUint8, biden.UnmarshalUint8)
}

func BenchmarkUnmarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	decodedUint16s = benchmarkUnmarshalSlice(b, nums, biden.Uint16Bytes, biden.MarshalUint16, biden.UnmarshalUint16)
}

func BenchmarkUnmarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	decodedUint32s = benchmarkUnmarshalSlice(b, nums, biden.Uint32Bytes, biden.MarshalUint32, biden.UnmarshalUint32)
}

func BenchmarkUnmarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	decodedUint64s = benchmarkUnmarshalSlice(b, nums, biden.Uint64Bytes, biden.MarshalUint64, biden.UnmarshalUint64)
}

func BenchmarkUnmarshalIntSlice(b *testing.B) {
	nums := makeInts()
	decodedInts = benchmarkUnmarshalSlice(b, nums, biden.IntBytes, biden.MarshalInt, biden.UnmarshalInt)
}

func BenchmarkUnmarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	decodedInt8s = benchmarkUnmarshalSlice(b, nums, biden.Int8Bytes, biden.MarshalInt8, biden.UnmarshalInt8)
}

func BenchmarkUnmarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	decodedInt16s = benchmarkUnmarshalSlice(b, nums, biden.Int16Bytes, biden.MarshalInt16, biden.UnmarshalInt16)
}

func BenchmarkUnmarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	decodedInt32s = benchmarkUnmarshalSlice(b, nums, biden.Int32Bytes, biden.MarshalInt32, biden.UnmarshalInt32)
}

func BenchmarkUnmarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	decodedInt64s = benchmarkUnmarshalSlice(b, nums, biden.Int64Bytes, biden.MarshalInt64, biden.UnmarshalInt64)
}

func BenchmarkUnmarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	decodedFloat32s = benchmarkUnmarshalSlice(b, nums, biden.Float32Bytes, biden.MarshalFloat32, biden.UnmarshalFloat32)
}

func BenchmarkUnmarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	decodedFloat64s = benchmarkUnmarshalSlice(b, nums, biden.Float64Bytes, biden.MarshalFloat64, biden.UnmarshalFloat64)
}

func BenchmarkUnmarshalStringSlice(b *testing.B) {
	var (
		strings = makeStrings()
		size    = biden.StringSliceBytes(strings)

		marshaler   = makeSliceMarshaler(biden.MarshalString)
		unmarshaler = makeSliceUnmarshaler(biden.UnmarshalString)

		encoded = biden.Marshal(strings, size, marshaler)
	)

	b.ResetTimer()

	var result []string
	for i := 0; i < b.N; i++ {
		result = biden.Unmarshal(encoded, unmarshaler)
	}
	decodedStrings = result
}

func benchmarkUnmarshalSlice[T any](b *testing.B, slice []T, itemSize int, itemMarshaler biden.MarshalFunc[T], itemUnmarshaler biden.UnmarshalFunc[T]) []T {
	var (
		size = biden.SliceBytes(slice, itemSize)

		marshaler   = makeSliceMarshaler(itemMarshaler)
		unmarshaler = makeSliceUnmarshaler(itemUnmarshaler)

		encoded = biden.Marshal(slice, size, marshaler)
	)

	b.ResetTimer()

	var result []T
	for i := 0; i < b.N; i++ {
		result = biden.Unmarshal(encoded, unmarshaler)
	}
	return result
}

func makeSliceUnmarshaler[T any](itemUnmarshaler biden.UnmarshalFunc[T]) biden.UnmarshalFunc[[]T] {
	return func(pos int, buf []byte) ([]T, int) {
		return biden.UnmarshalSlice(pos, buf, itemUnmarshaler)
	}
}
