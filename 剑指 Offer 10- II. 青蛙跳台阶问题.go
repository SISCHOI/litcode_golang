
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	a, b, c := 0, 1, 1
	for n > 1 {
		a = b
		b = c
		c = (a + b) % 1000000007
		n--
	}
	return c

}