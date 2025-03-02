package benchmark

import (
	"testing"

	biden "github.com/daspoet/biden/v1"
)

var encodeResult []byte

func BenchmarkMarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	encodeResult = benchmarkMarshalSlice(b, bools, biden.BoolBytes, biden.MarshalBool)
}

func BenchmarkMarshalUintSlice(b *testing.B) {
	nums := makeUints()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.UintBytes, biden.MarshalUint)
}

func BenchmarkMarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Uint8Bytes, biden.MarshalUint8)
}

func BenchmarkMarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Uint16Bytes, biden.MarshalUint16)
}

func BenchmarkMarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Uint32Bytes, biden.MarshalUint32)
}

func BenchmarkMarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Uint64Bytes, biden.MarshalUint64)
}

func BenchmarkMarshalIntSlice(b *testing.B) {
	nums := makeInts()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.IntBytes, biden.MarshalInt)
}

func BenchmarkMarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Int8Bytes, biden.MarshalInt8)
}

func BenchmarkMarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Int16Bytes, biden.MarshalInt16)
}

func BenchmarkMarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Int32Bytes, biden.MarshalInt32)
}

func BenchmarkMarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Int64Bytes, biden.MarshalInt64)
}

func BenchmarkMarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Float32Bytes, biden.MarshalFloat32)
}

func BenchmarkMarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	encodeResult = benchmarkMarshalSlice(b, nums, biden.Float64Bytes, biden.MarshalFloat64)
}

func BenchmarkMarshalStringSlice(b *testing.B) {
	var (
		strings = makeStrings()

		size      = biden.StringSliceBytes(strings)
		marshaler = makeSliceMarshaler(biden.MarshalString)
	)

	b.ResetTimer()

	var result []byte
	for i := 0; i < b.N; i++ {
		result = biden.Marshal(strings, size, marshaler)
	}
	encodeResult = result
}

func benchmarkMarshalSlice[T any](b *testing.B, slice []T, itemSize int, itemMarshaler biden.MarshalFunc[T]) []byte {
	var (
		size      = biden.SliceBytes(slice, itemSize)
		marshaler = makeSliceMarshaler(itemMarshaler)
	)

	b.ResetTimer()

	var result []byte
	for i := 0; i < b.N; i++ {
		result = biden.Marshal(slice, size, marshaler)
	}
	return result
}

func makeSliceMarshaler[T any](itemMarshaler biden.MarshalFunc[T]) biden.MarshalFunc[[]T] {
	return func(pos int, buf []byte, value []T) int {
		return biden.MarshalSlice(pos, buf, value, itemMarshaler)
	}
}
