func isSubsequence(s string, t string) bool {
    lastPosition := 0

    for _, rune := range s {
        hasMatch := false
        for j := lastPosition; j < len(t); j++ {
            if byte(rune) == t[j] {
                lastPosition = j + 1
                hasMatch = true
                break
            }
        }

        if !hasMatch {
            return false
        }
    }

    return true
}
