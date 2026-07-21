func canPlaceFlowers(flowerbed []int, n int) bool {
	current := n

   for i := 0; i < len(flowerbed); i++ {
		 var left int
		 var right int

		 if current == 0 {
			return true
		 }

		 if flowerbed[i] == 1 {
			continue
		 }

		 if i - 1 < 0 {
			 left = 0
		 } else {
			left = flowerbed[i - 1]
		 }

		 if i + 1 > len(flowerbed) - 1  {
			 right = 0
		 } else {
			right = flowerbed[i + 1]
		 }

		 if left != 1 && right != 1 {
			 current--
			 flowerbed[i] = 1
		 }
	 } 

	 return current == 0
}