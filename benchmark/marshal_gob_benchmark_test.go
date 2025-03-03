package benchmark

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func Benchmark_gob_MarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	benchmarkMarshalGobSlice(b, bools)
}

func Benchmark_gob_MarshalUintSlice(b *testing.B) {
	nums := makeUints()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalIntSlice(b *testing.B) {
	nums := makeInts()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	benchmarkMarshalGobSlice(b, nums)
}

func Benchmark_gob_MarshalStringSlice(b *testing.B) {
	strings := makeStrings()
	benchmarkMarshalGobSlice(b, strings)
}

func benchmarkMarshalGobSlice[T any](b *testing.B, slice []T) {
	for b.Loop() {
		var (
			buf     bytes.Buffer
			encoder = gob.NewEncoder(&buf)
		)

		_ = encoder.Encode(slice)
	}
}
