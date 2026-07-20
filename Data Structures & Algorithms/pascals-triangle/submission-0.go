func generate(numRows int) [][]int {
	row := 0	
	element := 1
	result := make([][]int, numRows)

	for row < numRows {
		result[row] = []int{}
		
		for i := range element {
			if i == 0 || i == element - 1 {
				result[row] = append(result[row], 1)
			} else {
				value := result[row-1][i - 1] + result[row-1][i]
				result[row] = append(result[row], value)
			}
		}

		row += 1
		element += 1
	}

	return result
}
