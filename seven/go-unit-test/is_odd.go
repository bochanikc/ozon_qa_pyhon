package unit

func IsOdd1(n int) bool {
	return n%2 == 0
}

func IsOdd2(n int) bool {
	return n&1 == 1
}
