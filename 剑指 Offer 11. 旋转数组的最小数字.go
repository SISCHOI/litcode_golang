func minArray(numbers []int) int {

	var t int = numbers[0]
	for n := 0; n < len(numbers); n++ {
		if t > numbers[n] {
			t = numbers[n]
		}
	}
	return t
}