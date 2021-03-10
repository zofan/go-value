package value

import "time"

func MinBool(a, b bool) bool {
	return a || b
}

func MinString(a, b string) string {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinTime(a, b time.Time) time.Time {
	if a.After(b) {
		return b
	}
	return a
}

func MinStringSlice(a, b []string) []string {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinIntSlice(a, b []int) []int {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinInt8Slice(a, b []int8) []int8 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinInt16Slice(a, b []int16) []int16 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinInt32Slice(a, b []int32) []int32 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinInt64Slice(a, b []int64) []int64 {
	if len(a) > len(b) {
		return b
	}
	return a
}
func MinUintSlice(a, b []uint) []uint {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinUint8Slice(a, b []uint8) []uint8 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinUint16Slice(a, b []uint16) []uint16 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinUint32Slice(a, b []uint32) []uint32 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinUint64Slice(a, b []uint64) []uint64 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinFloat32Slice(a, b []float32) []float32 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinFloat64Slice(a, b []float64) []float64 {
	if len(a) > len(b) {
		return b
	}
	return a
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinInt8(a, b int8) int8 {
	if a > b {
		return b
	}
	return a
}

func MinInt16(a, b int16) int16 {
	if a > b {
		return b
	}
	return a
}

func MinInt32(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

func MinInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func MinUint(a, b uint) uint {
	if a > b {
		return b
	}
	return a
}

func MinUint8(a, b uint8) uint8 {
	if a > b {
		return b
	}
	return a
}

func MinUint16(a, b uint16) uint16 {
	if a > b {
		return b
	}
	return a
}

func MinUint32(a, b uint32) uint32 {
	if a > b {
		return b
	}
	return a
}

func MinUint64(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

func MinFloat32(a, b float32) float32 {
	if a > b {
		return b
	}
	return a
}

func MinFloat64(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
