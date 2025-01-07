package biden

import (
	"math/rand"
	"strings"
	"testing"
)

const (
	sampleSize = 1_000_000
	stringSize = 1_000
)

var rng = rand.New(rand.NewSource(2306))

var encodeResult []byte

func BenchmarkMarshalBoolSlice(b *testing.B) {
	nums := makeBools()
	encodeResult = benchmarkMarshalSlice(b, nums, BoolBytes, MarshalBool)
}

func BenchmarkMarshalUintSlice(b *testing.B) {
	nums := makeUints()
	encodeResult = benchmarkMarshalSlice(b, nums, UintBytes, MarshalUint)
}

func BenchmarkMarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	encodeResult = benchmarkMarshalSlice(b, nums, Uint8Bytes, MarshalUint8)
}

func BenchmarkMarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	encodeResult = benchmarkMarshalSlice(b, nums, Uint16Bytes, MarshalUint16)
}

func BenchmarkMarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	encodeResult = benchmarkMarshalSlice(b, nums, Uint32Bytes, MarshalUint32)
}

func BenchmarkMarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	encodeResult = benchmarkMarshalSlice(b, nums, Uint64Bytes, MarshalUint64)
}

func BenchmarkMarshalIntSlice(b *testing.B) {
	nums := makeInts()
	encodeResult = benchmarkMarshalSlice(b, nums, IntBytes, MarshalInt)
}

func BenchmarkMarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	encodeResult = benchmarkMarshalSlice(b, nums, Int8Bytes, MarshalInt8)
}

func BenchmarkMarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	encodeResult = benchmarkMarshalSlice(b, nums, Int16Bytes, MarshalInt16)
}

func BenchmarkMarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	encodeResult = benchmarkMarshalSlice(b, nums, Int32Bytes, MarshalInt32)
}

func BenchmarkMarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	encodeResult = benchmarkMarshalSlice(b, nums, Int64Bytes, MarshalInt64)
}

func BenchmarkMarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	encodeResult = benchmarkMarshalSlice(b, nums, Float32Bytes, MarshalFloat32)
}

func BenchmarkMarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	encodeResult = benchmarkMarshalSlice(b, nums, Float64Bytes, MarshalFloat64)
}

func BenchmarkMarshalStringSlice(b *testing.B) {
	var (
		strings = makeStrings()

		size      = StringSliceBytes(strings)
		marshaler = makeSliceMarshaler(MarshalString)
	)

	b.ResetTimer()

	var result []byte
	for i := 0; i < b.N; i++ {
		result = Marshal(strings, size, marshaler)
	}
	encodeResult = result
}

func benchmarkMarshalSlice[T any](b *testing.B, slice []T, itemSize int, itemMarshaler MarshalFunc[T]) []byte {
	var (
		size      = SliceBytes(slice, itemSize)
		marshaler = makeSliceMarshaler(itemMarshaler)
	)

	b.ResetTimer()

	var result []byte
	for i := 0; i < b.N; i++ {
		result = Marshal(slice, size, marshaler)
	}
	return result
}

func makeBools() []bool {
	bools := make([]bool, sampleSize)
	for i := range sampleSize {
		bools[i] = rng.Intn(2) < 1
	}
	return bools
}

func makeUints() []uint {
	nums := make([]uint, sampleSize)
	for i := range sampleSize {
		nums[i] = uint(rng.Uint64())
	}
	return nums
}

func makeUint8s() []uint8 {
	nums := make([]uint8, sampleSize)
	for i := range sampleSize {
		nums[i] = uint8(rng.Uint32())
	}
	return nums
}

func makeUint16s() []uint16 {
	nums := make([]uint16, sampleSize)
	for i := range sampleSize {
		nums[i] = uint16(rng.Uint32())
	}
	return nums
}

func makeUint32s() []uint32 {
	nums := make([]uint32, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Uint32()
	}
	return nums
}

func makeUint64s() []uint64 {
	nums := make([]uint64, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Uint64()
	}
	return nums
}

func makeInts() []int {
	nums := make([]int, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Int()
	}
	return nums
}

func makeInt8s() []int8 {
	nums := make([]int8, sampleSize)
	for i := range sampleSize {
		nums[i] = int8(rng.Int31())
	}
	return nums
}

func makeInt16s() []int16 {
	nums := make([]int16, sampleSize)
	for i := range sampleSize {
		nums[i] = int16(rng.Int31())
	}
	return nums
}

func makeInt32s() []int32 {
	nums := make([]int32, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Int31()
	}
	return nums
}

func makeInt64s() []int64 {
	nums := make([]int64, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Int63()
	}
	return nums
}

func makeFloat32s() []float32 {
	nums := make([]float32, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Float32()
	}
	return nums
}

func makeFloat64s() []float64 {
	nums := make([]float64, sampleSize)
	for i := range sampleSize {
		nums[i] = rng.Float64()
	}
	return nums
}

func makeStrings() []string {
	items := make([]string, sampleSize)
	for i := range sampleSize {
		var b strings.Builder
		b.Grow(stringSize)
		for range stringSize {
			b.WriteRune(rune(rng.Uint32()))
		}
		items[i] = b.String()
	}
	return items
}
