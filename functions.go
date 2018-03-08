package listing

var factorial = make(map[int]int)

// Fac method
func Fac(n int) int {
	if len(factorial) < 2 {
		factorial[0], factorial[1] = 1, 1
	}
	if v, ok := factorial[n]; ok {
		return v
	}
	factorial[n] = n * Fac(n-1)
	return factorial[n]
}

// C method
func C(n, m int) (result int) {
	switch {
	case n == m:
		result = 1
	case m == 1 || m == 0:
		result = n
	default:
		result = C(n-1, m) + C(n-1, m-1)
	}
	return result
}

// P method
func P(n, m int) (result int) {
	return Fac(n) / Fac(n-m)
}

// H method
func H(n, m int) (result int) {
	return C(n+m-1, m)
}
