func countSeniors(details []string) int {
	result := 0
	for _, detail := range details {
		start := 11
		end := start + (len(detail) - 2 - start)
		ageStr := detail[start:end]

		age, _ := strconv.Atoi(ageStr)

		if age > 60 {
			result += 1
		}
	}

	return result
}