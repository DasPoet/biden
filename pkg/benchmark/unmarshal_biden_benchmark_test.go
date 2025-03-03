package benchmark

import (
	"github.com/daspoet/biden/pkg/biden"
	"testing"
)

func Benchmark_Biden_UnmarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	benchmarkUnmarshalSlice(b, bools, biden.BoolBytes, biden.MarshalBool, biden.UnmarshalBool)
}

func Benchmark_Biden_UnmarshalUintSlice(b *testing.B) {
	nums := makeUints()
	benchmarkUnmarshalSlice(b, nums, biden.UintBytes, biden.MarshalUint, biden.UnmarshalUint)
}

func Benchmark_Biden_UnmarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	benchmarkUnmarshalSlice(b, nums, biden.Uint8Bytes, biden.MarshalUint8, biden.UnmarshalUint8)
}

func Benchmark_Biden_UnmarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	benchmarkUnmarshalSlice(b, nums, biden.Uint16Bytes, biden.MarshalUint16, biden.UnmarshalUint16)
}

func Benchmark_Biden_UnmarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	benchmarkUnmarshalSlice(b, nums, biden.Uint32Bytes, biden.MarshalUint32, biden.UnmarshalUint32)
}

func Benchmark_Biden_UnmarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	benchmarkUnmarshalSlice(b, nums, biden.Uint64Bytes, biden.MarshalUint64, biden.UnmarshalUint64)
}

func Benchmark_Biden_UnmarshalIntSlice(b *testing.B) {
	nums := makeInts()
	benchmarkUnmarshalSlice(b, nums, biden.IntBytes, biden.MarshalInt, biden.UnmarshalInt)
}

func Benchmark_Biden_UnmarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	benchmarkUnmarshalSlice(b, nums, biden.Int8Bytes, biden.MarshalInt8, biden.UnmarshalInt8)
}

func Benchmark_Biden_UnmarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	benchmarkUnmarshalSlice(b, nums, biden.Int16Bytes, biden.MarshalInt16, biden.UnmarshalInt16)
}

func Benchmark_Biden_UnmarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	benchmarkUnmarshalSlice(b, nums, biden.Int32Bytes, biden.MarshalInt32, biden.UnmarshalInt32)
}

func Benchmark_Biden_UnmarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	benchmarkUnmarshalSlice(b, nums, biden.Int64Bytes, biden.MarshalInt64, biden.UnmarshalInt64)
}

func Benchmark_Biden_UnmarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	benchmarkUnmarshalSlice(b, nums, biden.Float32Bytes, biden.MarshalFloat32, biden.UnmarshalFloat32)
}

func Benchmark_Biden_UnmarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	benchmarkUnmarshalSlice(b, nums, biden.Float64Bytes, biden.MarshalFloat64, biden.UnmarshalFloat64)
}

func Benchmark_Biden_UnmarshalStringSlice(b *testing.B) {
	var (
		strings = makeStrings()

		size      = biden.StringSliceBytes(strings)
		marshaler = makeSliceMarshaler(biden.MarshalString)

		encoded = biden.Marshal(strings, size, marshaler)
	)

	for b.Loop() {
		unmarshaler := makeSliceUnmarshaler(biden.UnmarshalString)
		biden.Unmarshal(encoded, unmarshaler)
	}
}

func benchmarkUnmarshalSlice[T any](b *testing.B, slice []T, itemSize int, itemMarshaler biden.MarshalFunc[T], itemUnmarshaler biden.UnmarshalFunc[T]) {
	var (
		size      = biden.SliceBytes(slice, itemSize)
		marshaler = makeSliceMarshaler(itemMarshaler)

		encoded = biden.Marshal(slice, size, marshaler)
	)

	for b.Loop() {
		unmarshaler := makeSliceUnmarshaler(itemUnmarshaler)
		biden.Unmarshal(encoded, unmarshaler)
	}
}

func makeSliceUnmarshaler[T any](itemUnmarshaler biden.UnmarshalFunc[T]) biden.UnmarshalFunc[[]T] {
	return func(pos int, buf []byte) ([]T, int) {
		return biden.UnmarshalSlice(pos, buf, itemUnmarshaler)
	}
}
