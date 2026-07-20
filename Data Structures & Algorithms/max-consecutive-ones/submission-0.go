func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	continous := 0
	
	for _, num := range nums {
		if num == 1 {
			continous += 1
		} else {
			continous = 0
		}

		if continous > max {
			max = continous
		}
	}

	return  max
}
