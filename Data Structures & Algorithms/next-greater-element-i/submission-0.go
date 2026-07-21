func nextGreaterElement(nums1 []int, nums2 []int) []int {
	idxTrack := map[int]int{}
	result := []int{}

	for idx, num := range nums2 {
		idxTrack[num] = idx
	}

	for _, num := range nums1 {
		hasGreater := false
		for _, compare := range nums2[idxTrack[num]:] {
			if compare > num {
				result = append(result, compare)
				hasGreater = true
				break
			}
		}

		if !hasGreater {
			result = append(result, -1)
		}
	}

	return result
}
