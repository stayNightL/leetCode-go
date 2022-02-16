package offer
type CQueue struct {
	in_len int
	out_len int
	inStack []int
	outStack []int
}


func Constructor() CQueue {
	return CQueue{
		inStack: make([]int,10000),
		outStack: make([]int,10000),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.in_len++
	this.inStack[this.in_len] = value
}


func (this *CQueue) DeleteHead() int {
	if this.out_len==0 {
		for i := this.in_len; i > 0; i-- {
			this.out_len++
			this.outStack[this.out_len] = this.inStack[i]
		}
	}
	if this.out_len==0 {
		return -1
	}
	r := this.outStack[this.out_len]
	this.out_len--
	return r

}