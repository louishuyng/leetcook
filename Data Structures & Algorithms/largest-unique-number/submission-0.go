func largestUniqueNumber(nums []int) int {
    track := map[int]int{}
    result := -1

    for _, num := range nums {
        track[num] += 1
    }

    for key, value := range track {
        if value != 1 {
            continue
        }
        if key > result {
            result = key
        }
    }


    return result
}
