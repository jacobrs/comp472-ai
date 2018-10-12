package main

// Abs - Helper for returning the absolute value of an integer
func Abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}
