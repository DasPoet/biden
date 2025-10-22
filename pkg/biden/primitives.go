package biden

import (
	"math"

	"github.com/google/uuid"
)

const (
	BoolBytes = 1

	UintBytes   = 8
	Uint8Bytes  = 1
	Uint16Bytes = 2
	Uint32Bytes = 4
	Uint64Bytes = 8

	IntBytes   = 8
	Int8Bytes  = 1
	Int16Bytes = 2
	Int32Bytes = 4
	Int64Bytes = 8

	Float32Bytes = 4
	Float64Bytes = 8

	UUIDBytes = 16
)

func MarshalBool(pos int, buf []byte, b bool) int {
	if b {
		buf[pos] = 1
	} else {
		buf[pos] = 0
	}
	return pos + BoolBytes
}

func UnmarshalBool(pos int, buf []byte) (bool, int) {
	raw := buf[pos]
	return raw == 1, pos + BoolBytes
}

func MarshalUint8(pos int, buf []byte, n uint8) int {
	buf[pos] = byte(n)
	return pos + Uint8Bytes
}

func UnmarshalUint8(pos int, buf []byte) (uint8, int) {
	raw := buf[pos]
	return uint8(raw), pos + Uint8Bytes
}

func MarshalUint16(pos int, buf []byte, n uint16) int {
	window := buf[pos : pos+Uint16Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[1]

	window[0] = byte(n)
	window[1] = byte(n >> 8)

	return pos + Uint16Bytes
}

func UnmarshalUint16(pos int, buf []byte) (uint16, int) {
	window := buf[pos : pos+Uint16Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[1]

	lo := uint16(window[0])
	hi := uint16(window[1]) << 8

	return lo | hi, pos + Uint16Bytes
}

func MarshalUint32(pos int, buf []byte, n uint32) int {
	window := buf[pos : pos+Uint32Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[3]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)

	return pos + Uint32Bytes
}

func UnmarshalUint32(pos int, buf []byte) (uint32, int) {
	window := buf[pos : pos+Uint32Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[3]

	lo := uint32(window[0]) | uint32(window[1])<<8
	hi := uint32(window[2])<<16 | uint32(window[3])<<24

	return lo | hi, pos + Uint32Bytes
}

func MarshalUint64(pos int, buf []byte, n uint64) int {
	window := buf[pos : pos+Uint64Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[7]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)
	window[4] = byte(n >> 32)
	window[5] = byte(n >> 40)
	window[6] = byte(n >> 48)
	window[7] = byte(n >> 56)

	return pos + Uint64Bytes
}

func UnmarshalUint64(pos int, buf []byte) (uint64, int) {
	window := buf[pos : pos+Uint64Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[7]

	lo := uint64(window[0]) | uint64(window[1])<<8 | uint64(window[2])<<16 | uint64(window[3])<<24
	hi := uint64(window[4])<<32 | uint64(window[5])<<40 | uint64(window[6])<<48 | uint64(window[7])<<56

	return lo | hi, pos + Uint64Bytes
}

func MarshalUint(pos int, buf []byte, n uint) int {
	window := buf[pos : pos+UintBytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[7]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)
	window[4] = byte(n >> 32)
	window[5] = byte(n >> 40)
	window[6] = byte(n >> 48)
	window[7] = byte(n >> 56)

	return pos + UintBytes
}

func UnmarshalUint(pos int, buf []byte) (uint, int) {
	window := buf[pos : pos+Uint64Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[7]

	lo := uint(window[0]) | uint(window[1])<<8 | uint(window[2])<<16 | uint(window[3])<<24
	hi := uint(window[4])<<32 | uint(window[5])<<40 | uint(window[6])<<48 | uint(window[7])<<56

	return lo | hi, pos + UintBytes
}

func MarshalInt8(pos int, buf []byte, n int8) int {
	buf[pos] = byte(n)
	return pos + Int8Bytes
}

func UnmarshalInt8(pos int, buf []byte) (int8, int) {
	raw := buf[pos]
	return int8(raw), pos + Int8Bytes
}

func MarshalInt16(pos int, buf []byte, n int16) int {
	window := buf[pos : pos+Int16Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[1]

	window[0] = byte(n)
	window[1] = byte(n >> 8)

	return pos + Int16Bytes
}

func UnmarshalInt16(pos int, buf []byte) (int16, int) {
	window := buf[pos : pos+Int16Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[1]

	lo := int16(window[0])
	hi := int16(window[1]) << 8

	return lo | hi, pos + Int16Bytes
}

func MarshalInt32(pos int, buf []byte, n int32) int {
	window := buf[pos : pos+Int32Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[3]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)

	return pos + Int32Bytes
}

func UnmarshalInt32(pos int, buf []byte) (int32, int) {
	window := buf[pos : pos+Int32Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[3]

	lo := int32(window[0]) | int32(window[1])<<8
	hi := int32(window[2])<<16 | int32(window[3])<<24

	return lo | hi, pos + Int32Bytes
}

func MarshalInt64(pos int, buf []byte, n int64) int {
	window := buf[pos : pos+Int64Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[7]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)
	window[4] = byte(n >> 32)
	window[5] = byte(n >> 40)
	window[6] = byte(n >> 48)
	window[7] = byte(n >> 56)

	return pos + Int64Bytes
}

func UnmarshalInt64(pos int, buf []byte) (int64, int) {
	window := buf[pos : pos+Int64Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[7]

	lo := int64(window[0]) | int64(window[1])<<8 | int64(window[2])<<16 | int64(window[3])<<24
	hi := int64(window[4])<<32 | int64(window[5])<<40 | int64(window[6])<<48 | int64(window[7])<<56

	return lo | hi, pos + Int64Bytes
}

func MarshalFloat32(pos int, buf []byte, n float32) int {
	bits := math.Float32bits(n)

	window := buf[pos : pos+Float32Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[3]

	window[0] = byte(bits)
	window[1] = byte(bits >> 8)
	window[2] = byte(bits >> 16)
	window[3] = byte(bits >> 24)

	return pos + Float32Bytes
}

func UnmarshalFloat32(pos int, buf []byte) (float32, int) {
	window := buf[pos : pos+Float32Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[3]

	lo := uint32(window[0]) | uint32(window[1])<<8
	hi := uint32(window[2])<<16 | uint32(window[3])<<24

	return math.Float32frombits(lo | hi), pos + Float32Bytes
}

func MarshalFloat64(pos int, buf []byte, n float64) int {
	bits := math.Float64bits(n)

	window := buf[pos : pos+Float64Bytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[7]

	window[0] = byte(bits)
	window[1] = byte(bits >> 8)
	window[2] = byte(bits >> 16)
	window[3] = byte(bits >> 24)
	window[4] = byte(bits >> 32)
	window[5] = byte(bits >> 40)
	window[6] = byte(bits >> 48)
	window[7] = byte(bits >> 56)

	return pos + Float64Bytes
}

func UnmarshalFloat64(pos int, buf []byte) (float64, int) {
	window := buf[pos : pos+Int64Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[7]

	lo := uint64(window[0]) | uint64(window[1])<<8 | uint64(window[2])<<16 | uint64(window[3])<<24
	hi := uint64(window[4])<<32 | uint64(window[5])<<40 | uint64(window[6])<<48 | uint64(window[7])<<56

	return math.Float64frombits(lo | hi), pos + Int64Bytes
}

func MarshalInt(pos int, buf []byte, n int) int {
	window := buf[pos : pos+IntBytes]

	// early bounds check to guarantee safety of
	// writes below and to increase performance
	_ = window[7]

	window[0] = byte(n)
	window[1] = byte(n >> 8)
	window[2] = byte(n >> 16)
	window[3] = byte(n >> 24)
	window[4] = byte(n >> 32)
	window[5] = byte(n >> 40)
	window[6] = byte(n >> 48)
	window[7] = byte(n >> 56)

	return pos + IntBytes
}

func UnmarshalInt(pos int, buf []byte) (int, int) {
	window := buf[pos : pos+Int64Bytes]

	// early bounds check to guarantee safety of
	// reads below and to increase performance
	_ = window[7]

	lo := int(window[0]) | int(window[1])<<8 | int(window[2])<<16 | int(window[3])<<24
	hi := int(window[4])<<32 | int(window[5])<<40 | int(window[6])<<48 | int(window[7])<<56

	return lo | hi, pos + IntBytes
}

func StringBytes(s string) int {
	return IntBytes + len(s)
}

func MarshalString(pos int, buf []byte, s string) int {
	pos = MarshalInt(pos, buf, len(s))
	copy(buf[pos:], s)

	return pos + len(s)
}

func UnmarshalString(pos int, buf []byte) (string, int) {
	len, pos := UnmarshalInt(pos, buf)

	return string(buf[pos : pos+len]), pos + len
}

func MarshalUUID(pos int, buf []byte, id uuid.UUID) int {
	copy(buf[pos:pos+UUIDBytes], id[:])
	return pos + UUIDBytes
}

func UnmarshalUUID(pos int, buf []byte) (uuid.UUID, int) {
	var id uuid.UUID
	copy(id[:], buf[pos:pos+UUIDBytes])

	return id, pos + UUIDBytes
}
