func maxDifference(s string) int {
	frequency := map[byte]int{}

	for _, char := range s {
		frequency[byte(char)]++
	}

	maxOdd := 0
	minEven := 100

	for _, val := range frequency {
		if val % 2 == 0 && val < minEven {
			minEven = val
		}

		if val % 2 != 0 && val > maxOdd {
			maxOdd = val
		}
	}

	return  maxOdd - minEven
}
