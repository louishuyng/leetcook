func isAnagram(s string, t string) bool {
    dp1 := map[rune]int{}
    dp2 := map[rune]int{}

    for _, r := range s {
        if dp1[r] == 0 {
            dp1[r] = 1
        } else {
            dp1[r] += 1
        }
    }

    for _, r := range t {
        if dp2[r] == 0 {
            dp2[r] = 1
        } else {
            dp2[r] += 1
        }
    }

    if len(dp1) != len(dp2) {
        return false
    }

    for key, value := range dp1 {
        if dp2[key] != value {
            return false
        }
    }

    return true
}
