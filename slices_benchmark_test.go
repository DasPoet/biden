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

var (
	encodeResult []byte

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

	decodedBools   []bool
	decodedStrings []string
)

func BenchmarkMarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	encodeResult = benchmarkMarshalSlice(b, bools, BoolBytes, MarshalBool)
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

func BenchmarkUnmarshalBoolSlice(b *testing.B) {
	bools := makeBools()
	decodedBools = benchmarkUnmarshalSlice(b, bools, BoolBytes, MarshalBool, UnmarshalBool)
}

func BenchmarkUnmarshalUintSlice(b *testing.B) {
	nums := makeUints()
	decodedUints = benchmarkUnmarshalSlice(b, nums, UintBytes, MarshalUint, UnmarshalUint)
}

func BenchmarkUnmarshalUint8Slice(b *testing.B) {
	nums := makeUint8s()
	decodedUint8s = benchmarkUnmarshalSlice(b, nums, Uint8Bytes, MarshalUint8, UnmarshalUint8)
}

func BenchmarkUnmarshalUint16Slice(b *testing.B) {
	nums := makeUint16s()
	decodedUint16s = benchmarkUnmarshalSlice(b, nums, Uint16Bytes, MarshalUint16, UnmarshalUint16)
}

func BenchmarkUnmarshalUint32Slice(b *testing.B) {
	nums := makeUint32s()
	decodedUint32s = benchmarkUnmarshalSlice(b, nums, Uint32Bytes, MarshalUint32, UnmarshalUint32)
}

func BenchmarkUnmarshalUint64Slice(b *testing.B) {
	nums := makeUint64s()
	decodedUint64s = benchmarkUnmarshalSlice(b, nums, Uint64Bytes, MarshalUint64, UnmarshalUint64)
}

func BenchmarkUnmarshalIntSlice(b *testing.B) {
	nums := makeInts()
	decodedInts = benchmarkUnmarshalSlice(b, nums, IntBytes, MarshalInt, UnmarshalInt)
}

func BenchmarkUnmarshalInt8Slice(b *testing.B) {
	nums := makeInt8s()
	decodedInt8s = benchmarkUnmarshalSlice(b, nums, Int8Bytes, MarshalInt8, UnmarshalInt8)
}

func BenchmarkUnmarshalInt16Slice(b *testing.B) {
	nums := makeInt16s()
	decodedInt16s = benchmarkUnmarshalSlice(b, nums, Int16Bytes, MarshalInt16, UnmarshalInt16)
}

func BenchmarkUnmarshalInt32Slice(b *testing.B) {
	nums := makeInt32s()
	decodedInt32s = benchmarkUnmarshalSlice(b, nums, Int32Bytes, MarshalInt32, UnmarshalInt32)
}

func BenchmarkUnmarshalInt64Slice(b *testing.B) {
	nums := makeInt64s()
	decodedInt64s = benchmarkUnmarshalSlice(b, nums, Int64Bytes, MarshalInt64, UnmarshalInt64)
}

func BenchmarkUnmarshalFloat32Slice(b *testing.B) {
	nums := makeFloat32s()
	decodedFloat32s = benchmarkUnmarshalSlice(b, nums, Float32Bytes, MarshalFloat32, UnmarshalFloat32)
}

func BenchmarkUnmarshalFloat64Slice(b *testing.B) {
	nums := makeFloat64s()
	decodedFloat64s = benchmarkUnmarshalSlice(b, nums, Float64Bytes, MarshalFloat64, UnmarshalFloat64)
}

func BenchmarkUnmarshalStringSlice(b *testing.B) {
	var (
		strings = makeStrings()
		size    = StringSliceBytes(strings)

		marshaler   = makeSliceMarshaler(MarshalString)
		unmarshaler = makeSliceUnmarshaler(UnmarshalString)

		encoded = Marshal(strings, size, marshaler)
	)

	b.ResetTimer()

	var result []string
	for i := 0; i < b.N; i++ {
		result = Unmarshal(encoded, unmarshaler)
	}
	decodedStrings = result
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

func benchmarkUnmarshalSlice[T any](b *testing.B, slice []T, itemSize int, itemMarshaler MarshalFunc[T], itemUnmarshaler UnmarshalFunc[T]) []T {
	var (
		size = SliceBytes(slice, itemSize)

		marshaler   = makeSliceMarshaler(itemMarshaler)
		unmarshaler = makeSliceUnmarshaler(itemUnmarshaler)

		encoded = Marshal(slice, size, marshaler)
	)

	b.ResetTimer()

	var result []T
	for i := 0; i < b.N; i++ {
		result = Unmarshal(encoded, unmarshaler)
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
