package math

// Dim returns the maximum of a-b or 0 for ints
func DimInt(a, b int) int {
	return MaxInt(a-b, 0)
}

// Dim returns the maximum of a-b or 0 for int64s
func DimInt64(a, b int64) int64 {
	return MaxInt64(a-b, 0)
}

// Returns the maximum of a and b for ints
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the minimum of a and b for ints
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns the maximum of a and b for int64s
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Returns the minimum of a and b for int64s
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
