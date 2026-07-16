func calculateTime(keyboard string, word string) int {
   keyPosition := map[rune]int{}

   for idx, key := range keyboard {
        keyPosition[key] =  idx
   }

    lastPosition := 0
    sum := 0
   for _, key := range word {
        sum += int(math.Abs(float64(lastPosition - keyPosition[key])))
        lastPosition = keyPosition[key]

   }

   return sum
}
