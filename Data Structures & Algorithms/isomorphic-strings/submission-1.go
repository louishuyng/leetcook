func isIsomorphic(s string, t string) bool {
		mapST := make(map[byte]byte)
    mapTS := make(map[byte]byte)
		
    for i := 0; i < len(s); i++ {
        c1, c2 := s[i], t[i]
        if mapST[c1] != 0 && mapST[c1] != c2 {
            return false
        }
        if mapTS[c2] != 0 && mapTS[c2] != c1 {
            return false
        }
        mapST[c1] = c2
        mapTS[c2] = c1
    }
    return true
}
