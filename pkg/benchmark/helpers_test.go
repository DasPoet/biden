package benchmark

import (
	"math/rand"
	"strings"
)

const (
	sampleSize = 1_000_000
	stringSize = 1_000
)

var rng = rand.New(rand.NewSource(2306))

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
