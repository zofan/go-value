package value

import "time"

func MaxBool(a, b bool) bool {
	return a || b
}

func MaxString(a, b string) string {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return b
	}
	return a
}

func MaxStringSlice(a, b []string) []string {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxIntSlice(a, b []int) []int {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxInt8Slice(a, b []int8) []int8 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxInt16Slice(a, b []int16) []int16 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxInt32Slice(a, b []int32) []int32 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxInt64Slice(a, b []int64) []int64 {
	if len(a) < len(b) {
		return b
	}
	return a
}
func MaxUintSlice(a, b []uint) []uint {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxUint8Slice(a, b []uint8) []uint8 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxUint16Slice(a, b []uint16) []uint16 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxUint32Slice(a, b []uint32) []uint32 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxUint64Slice(a, b []uint64) []uint64 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxFloat32Slice(a, b []float32) []float32 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxFloat64Slice(a, b []float64) []float64 {
	if len(a) < len(b) {
		return b
	}
	return a
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxInt8(a, b int8) int8 {
	if a < b {
		return b
	}
	return a
}

func MaxInt16(a, b int16) int16 {
	if a < b {
		return b
	}
	return a
}

func MaxInt32(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func MaxUint(a, b uint) uint {
	if a < b {
		return b
	}
	return a
}

func MaxUint8(a, b uint8) uint8 {
	if a < b {
		return b
	}
	return a
}

func MaxUint16(a, b uint16) uint16 {
	if a < b {
		return b
	}
	return a
}

func MaxUint32(a, b uint32) uint32 {
	if a < b {
		return b
	}
	return a
}

func MaxUint64(a, b uint64) uint64 {
	if a < b {
		return b
	}
	return a
}

func MaxFloat32(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func MaxFloat64(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}
