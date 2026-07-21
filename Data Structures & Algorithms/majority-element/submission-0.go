func majorityElement(nums []int) int {
	currMax := 0
	result := 0
	track := map[int]int{}

	for _, num := range nums {
		track[num]++
		if track[num] > currMax {
			currMax = track[num]
			result = num
		}
	}

	return result
}
