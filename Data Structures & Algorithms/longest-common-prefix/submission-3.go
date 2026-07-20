func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}

	idx := 0
	current := strs[0][idx]
	result := []byte{}

	for {
		noMatch := false
		for _, str := range strs[1:] {
			if idx > len(str) - 1 {
				noMatch = true
				break
			}

			if str[idx] != current {
				noMatch = true
				break
			}
		}

		if noMatch {
			break
		} else {
			result = append(result, current)
		}

		if idx + 1 == len(strs[0]) {
			break
		} else {
			idx += 1
			current = strs[0][idx]
		}
	}

	return string(result)
}
