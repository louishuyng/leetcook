func hasDuplicate(nums []int) bool {
    dp := map[int]bool{}

    for _, num := range nums {
        if !dp[num] {
            dp[num] = true
        } else {
            return true
        }
    }
    
    return false
}
