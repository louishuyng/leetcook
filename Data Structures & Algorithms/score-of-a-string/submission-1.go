func scoreOfString(s string) int {
    i, j, sum := 0, 1, 0

    for j < len(s)  {
       diff := int(s[i]) - int(s[j])
       sum += int(math.Abs(float64(diff)))

       i += 1
       j += 1
    }

    return sum
}