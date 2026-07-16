func anagramMappings(nums1 []int, nums2 []int) []int {
	result := []int{}
	hash := map[int]int{}

	for idx, num := range nums2 {
		hash[num] = idx
	}

	for _, num := range nums1 {
		result = append(result, hash[num])
	}

	return result
}
