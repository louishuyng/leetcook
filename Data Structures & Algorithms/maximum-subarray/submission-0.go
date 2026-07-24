func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := 0

	for _, num := range nums {
		currSum = int(math.Max(float64(currSum), float64(0)))
		currSum += num

		maxSum = int(math.Max(float64(maxSum), float64(currSum)))
	}

	return maxSum
}
