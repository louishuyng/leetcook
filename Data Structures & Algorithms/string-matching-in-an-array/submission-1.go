func stringMatching(words []string) []string {
	var result []string
	seen := map[string]bool{}

	for i := 0; i < len(words) - 1; i++ {
		for j := i + 1; j < len(words); j++ {
			if len(words[i]) > len(words[j]) {
				if seen[words[j]] {
					continue
				}
				
				if isSubString(words[i], words[j]) {
					result = append(result, words[j])
					seen[words[j]] = true
				}
			} else {
				if seen[words[i]] {
					continue
				}

				if isSubString(words[j], words[i]) {
					result = append(result, words[i])
					seen[words[i]] = true
				}
			}
		}
	}

	return result
}

// Asume that len(word1) > len(word2)
func isSubString(word1, word2 string) bool {
	currentMatch := 0

	for _, word := range word1 {
		if currentMatch == len(word2) {
			return true
		}

		if byte(word) == word2[currentMatch] {
			currentMatch += 1
		} else {
			currentMatch = 0
		}
	}

	return currentMatch == len(word2)
}