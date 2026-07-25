func numOfSubarrays(arr []int, k int, threshold int) int {
    res := 0
    curSum := 0

    for i := 0; i < k-1; i++ {
        curSum += arr[i]
    }

    for L := 0; L <= len(arr)-k; L++ {
        curSum += arr[L+k-1]
        if curSum/k >= threshold {
            res++
        }
        curSum -= arr[L]
    }
    return res
}