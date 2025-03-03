package benchmark

import (
	"github.com/deneonet/benc/std"
	"testing"
)

func Benchmark_Benc_UnmarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	benchmarkUnmarshalBencSlice(b, bools, bstd.SizeBool, bstd.MarshalBool, bstd.UnmarshalBool)
}

func Benchmark_Benc_UnmarshalUintSlice(b *testing.B) {
	nums := makeUints()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeUint, bstd.MarshalUint, bstd.UnmarshalUint)
}

func Benchmark_Benc_UnmarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeByte, bstd.MarshalByte, bstd.UnmarshalByte)
}

func Benchmark_Benc_UnmarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeUint16, bstd.MarshalUint16, bstd.UnmarshalUint16)
}

func Benchmark_Benc_UnmarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeUint32, bstd.MarshalUint32, bstd.UnmarshalUint32)
}

func Benchmark_Benc_UnmarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeUint64, bstd.MarshalUint64, bstd.UnmarshalUint64)
}

func Benchmark_Benc_UnmarshalIntSlice(b *testing.B) {
	nums := makeInts()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeInt, bstd.MarshalInt, bstd.UnmarshalInt)
}

func Benchmark_Benc_UnmarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeInt16, bstd.MarshalInt16, bstd.UnmarshalInt16)
}

func Benchmark_Benc_UnmarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeInt32, bstd.MarshalInt32, bstd.UnmarshalInt32)
}

func Benchmark_Benc_UnmarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeInt64, bstd.MarshalInt64, bstd.UnmarshalInt64)
}

func Benchmark_Benc_UnmarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeFloat32, bstd.MarshalFloat32, bstd.UnmarshalFloat32)
}

func Benchmark_Benc_UnmarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	benchmarkUnmarshalBencSlice(b, nums, bstd.SizeFloat64, bstd.MarshalFloat64, bstd.UnmarshalFloat64)
}

func Benchmark_Benc_UnmarshalStringSlice(b *testing.B) {
	strings := makeStrings()
	benchmarkUnmarshalBencSlice(b, strings, bstd.SizeString, bstd.MarshalString, bstd.UnmarshalString)
}

func benchmarkUnmarshalBencSlice[T any](b *testing.B, slice []T, sizer any, marshaler bstd.MarshalFunc[T], unmarshaler any) {
	var (
		size = bstd.SizeSlice(slice, sizer)
		buf  = make([]byte, size)
	)

	bstd.MarshalSlice(0, buf, slice, marshaler)

	for b.Loop() {
		bstd.UnmarshalSlice[T](0, buf, unmarshaler)
	}
}
