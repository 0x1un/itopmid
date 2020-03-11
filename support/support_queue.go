package support

type RetryQueue []ResponseContent

func (self *RetryQueue) Push(respContent interface{}) {
	*self = append(*self, (ResponseContent)(respContent.(ResponseContent)))
}

func (self *RetryQueue) Pop() interface{} {
	if n := len(*self); n == 0 {
		return nil
	} else {
		v := (*self)[n-1]
		*self = (*self)[:n-1]
		return v
	}
}

func (self *RetryQueue) Len() int {
	return len(*self)
}
