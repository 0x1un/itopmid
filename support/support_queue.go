package support

type Queue []ResponseContent

func (self *Queue) Push(respContent interface{}) {
	*self = append(*self, (ResponseContent)(respContent.(ResponseContent)))
}

func (self *Queue) Pop() interface{} {
	if n := len(*self); n == 0 {
		return nil
	} else {
		v := (*self)[n-1]
		*self = (*self)[:n-1]
		return v
	}
}

func (self *Queue) Len() int {
	return len(*self)
}
