package benchmark

import (
	"testing"

	"github.com/daspoet/biden/pkg/biden"
)

func Benchmark_Biden_MarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	benchmarkMarshalSlice(b, bools, biden.BoolBytes, biden.MarshalBool)
}

func Benchmark_Biden_MarshalUintSlice(b *testing.B) {
	nums := makeUints()
	benchmarkMarshalSlice(b, nums, biden.UintBytes, biden.MarshalUint)
}

func Benchmark_Biden_MarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	benchmarkMarshalSlice(b, nums, biden.Uint8Bytes, biden.MarshalUint8)
}

func Benchmark_Biden_MarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	benchmarkMarshalSlice(b, nums, biden.Uint16Bytes, biden.MarshalUint16)
}

func Benchmark_Biden_MarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	benchmarkMarshalSlice(b, nums, biden.Uint32Bytes, biden.MarshalUint32)
}

func Benchmark_Biden_MarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	benchmarkMarshalSlice(b, nums, biden.Uint64Bytes, biden.MarshalUint64)
}

func Benchmark_Biden_MarshalIntSlice(b *testing.B) {
	nums := makeInts()
	benchmarkMarshalSlice(b, nums, biden.IntBytes, biden.MarshalInt)
}

func Benchmark_Biden_MarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	benchmarkMarshalSlice(b, nums, biden.Int8Bytes, biden.MarshalInt8)
}

func Benchmark_Biden_MarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	benchmarkMarshalSlice(b, nums, biden.Int16Bytes, biden.MarshalInt16)
}

func Benchmark_Biden_MarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	benchmarkMarshalSlice(b, nums, biden.Int32Bytes, biden.MarshalInt32)
}

func Benchmark_Biden_MarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	benchmarkMarshalSlice(b, nums, biden.Int64Bytes, biden.MarshalInt64)
}

func Benchmark_Biden_MarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	benchmarkMarshalSlice(b, nums, biden.Float32Bytes, biden.MarshalFloat32)
}

func Benchmark_Biden_MarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	benchmarkMarshalSlice(b, nums, biden.Float64Bytes, biden.MarshalFloat64)
}

func Benchmark_Biden_MarshalStringSlice(b *testing.B) {
	for b.Loop() {
		var (
			strings = makeStrings()

			size      = biden.StringSliceBytes(strings)
			marshaler = makeSliceMarshaler(biden.MarshalString)
		)

		biden.Marshal(strings, size, marshaler)
	}
}

func benchmarkMarshalSlice[T any](b *testing.B, slice []T, itemSize int, itemMarshaler biden.MarshalFunc[T]) {
	for b.Loop() {
		var (
			size      = biden.SliceBytes(slice, itemSize)
			marshaler = makeSliceMarshaler(itemMarshaler)
		)

		biden.Marshal(slice, size, marshaler)
	}
}

func makeSliceMarshaler[T any](itemMarshaler biden.MarshalFunc[T]) biden.MarshalFunc[[]T] {
	return func(pos int, buf []byte, value []T) int {
		return biden.MarshalSlice(pos, buf, value, itemMarshaler)
	}
}
