func lengthOfLastWord(s string) int {
	count := 0
	j := len(s) - 1

	if j == 0 && s[j] != ' ' {
		return 1
	}

	for s[j] == ' ' {
		j--
	}
	
	for j != 0 {
		if s[j] != ' ' {
			count += 1
			j--
		} else {
			break
		}
	}

	return count
}
