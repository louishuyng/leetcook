func appendCharacters(s string, t string) int {
    i, j := 0, 0

    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
           i += 1 
           j += 1
        } else {
           i += 1
        }
    }

    return len(t) - j
}