package benchmark

import (
	"github.com/deneonet/benc/std"
	"testing"
)

func Benchmark_Benc_MarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	benchmarkMarshalBencSlice(b, bools, bstd.SizeBool, bstd.MarshalBool)
}

func Benchmark_Benc_MarshalUintSlice(b *testing.B) {
	nums := makeUints()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeUint, bstd.MarshalUint)
}

func Benchmark_Benc_MarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeByte, bstd.MarshalByte)
}

func Benchmark_Benc_MarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeUint16, bstd.MarshalUint16)
}

func Benchmark_Benc_MarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeUint32, bstd.MarshalUint32)
}

func Benchmark_Benc_MarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeUint64, bstd.MarshalUint64)
}

func Benchmark_Benc_MarshalIntSlice(b *testing.B) {
	nums := makeInts()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeInt, bstd.MarshalInt)
}

func Benchmark_Benc_MarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeInt16, bstd.MarshalInt16)
}

func Benchmark_Benc_MarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeInt32, bstd.MarshalInt32)
}

func Benchmark_Benc_MarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeInt64, bstd.MarshalInt64)
}

func Benchmark_Benc_MarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeFloat32, bstd.MarshalFloat32)
}

func Benchmark_Benc_MarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	benchmarkMarshalBencSlice(b, nums, bstd.SizeFloat64, bstd.MarshalFloat64)
}

func Benchmark_Benc_MarshalStringSlice(b *testing.B) {
	strings := makeStrings()
	benchmarkMarshalBencSlice(b, strings, bstd.SizeString, bstd.MarshalString)
}

func benchmarkMarshalBencSlice[T any](b *testing.B, slice []T, sizer any, marshaler bstd.MarshalFunc[T]) {
	for b.Loop() {
		var (
			size = bstd.SizeSlice(slice, sizer)
			buf  = make([]byte, size)
		)

		_ = bstd.MarshalSlice(0, buf, slice, marshaler)
	}
}
