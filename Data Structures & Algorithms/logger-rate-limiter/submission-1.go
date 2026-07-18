type Logger struct {
	timeTrack map[string]int
}

func Constructor() Logger {
	return Logger{
		timeTrack: map[string]int{},
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if this.timeTrack[message] == 0 || this.timeTrack[message] <= timestamp {
		this.timeTrack[message] = timestamp + 10
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
