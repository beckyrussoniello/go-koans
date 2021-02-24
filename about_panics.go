package go_koans

func divideFourBy(i int) int {
	defaultAnswer := 2
	if i == 0 {
		return defaultAnswer
	}
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
