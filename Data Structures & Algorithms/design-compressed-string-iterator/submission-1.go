type StringIterator struct {
    res string
    ptr int
    num int
    ch  byte
}

func Constructor(compressedString string) StringIterator {
    return StringIterator{res: compressedString, ptr: 0, num: 0, ch: ' '}
}

func (this *StringIterator) Next() byte {
    if !this.HasNext() {
        return ' '
    }

    if this.num == 0 {
        this.ch = this.res[this.ptr]
        this.ptr++
        for this.ptr < len(this.res) && this.res[this.ptr] >= '0' && this.res[this.ptr] <= '9' {
            this.num = this.num*10 + int(this.res[this.ptr]-'0')
            this.ptr++
        }
    }

    this.num--
    return this.ch
}

func (this *StringIterator) HasNext() bool {
    return this.ptr != len(this.res) || this.num != 0
}