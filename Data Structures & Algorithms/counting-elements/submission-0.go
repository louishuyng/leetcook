func countElements(arr []int) int {
    hash := map[int]bool{}

    for _, val := range arr {
        hash[val] = true
    }

    sum := 0
    for  _, val := range arr {
        if hash[val + 1] {
            sum ++
        }
    }

    return sum
}