func stringShift(s string, shift [][]int) string {
    left := 0
    right := 0
    decision := ""
    step := 0

    for _, shf := range shift {
        if shf[0] == 0 {
            left += shf[1]
        } else {
            right += shf[1]
        }
    }
    if left > right {
        decision = "left"
        step = left - right
    } else {
        decision = "right"
        step = right - left
    }

    n := len(s)
    result := make([]rune, n)
    for idx, word := range s {
        var newPosition int

        if decision == "left" {
            newPosition = idx - step
        } else {
            newPosition = idx + step
        }
        newPosition = ((newPosition % n) + n) % n

        result[newPosition] = word
    }

    return string(result)
}