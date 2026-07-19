type MovingAverage struct {
 window []int
 total int
 size int
}

func Constructor(size int) *MovingAverage {
	return &MovingAverage{
		window: []int{},
		total: 0,
		size: size,
	}

}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.window) == this.size {
		removed := this.window[0]
		this.window = this.window[1:]
		this.total -= removed
	}

	this.total += val
	this.window = append(this.window, val)

	return float64(this.total) / float64(len(this.window))
}
