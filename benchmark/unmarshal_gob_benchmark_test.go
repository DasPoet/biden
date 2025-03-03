package benchmark

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_gob_UnmarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	benchmarkUnmarshalGobSlice(b, bools)
}

func Benchmark_gob_UnmarshalUintSlice(b *testing.B) {
	nums := makeUints()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalIntSlice(b *testing.B) {
	nums := makeInts()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	benchmarkUnmarshalGobSlice(b, nums)
}

func Benchmark_gob_UnmarshalStringSlice(b *testing.B) {
	strings := makeStrings()
	benchmarkUnmarshalGobSlice(b, strings)
}

func benchmarkUnmarshalGobSlice[T any](b *testing.B, slice []T) {
	var (
		writeBuf bytes.Buffer
		encoder  = gob.NewEncoder(&writeBuf)
	)

	err := encoder.Encode(slice)
	require.NoError(b, err)

	for b.Loop() {
		var (
			reader  = bytes.NewReader(writeBuf.Bytes())
			decoder = gob.NewDecoder(reader)
		)

		var decoded []T
		_ = decoder.Decode(&decoded)
	}
}
