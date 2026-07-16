func confusingNumber(n int) bool {
	keys := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}

	sum := 0
	curr := n
	for i := totalNumbers(n); i != 0; i-- {
		num := curr / int(math.Pow(10, float64(i - 1)))
		curr = curr % int(math.Pow(10, float64(i - 1)))

		rotated, ok := keys[num]

		if !ok {
			return false
		}

		sum += rotated * int(math.Pow(10, float64(totalNumbers(n) - i)))
	}

	if sum != n {
		return true
	}

	return false
}

func totalNumbers(n int) int {
	curr := n
	count := 0
	for  curr != 0 {
		curr /= 10
		count += 1
	}

	return count
}
