package support

type Queue []interface{}

type TicketQueue map[string]string

func (self TicketQueue) Set(k, v string) {
	self[k] = v
}

func (self TicketQueue) Del(k string) {
	delete(self, k)
}

func (self TicketQueue) Self() map[string]string {
	return self
}

func (self TicketQueue) Get(key string) string {
	if v, ok := self[key]; !ok {
		return ""
	} else {
		return v
	}
}

func (self *Queue) Push(respContent interface{}) {
	switch respContent.(type) {
	case string:
		*self = append(*self, respContent.(string))
	case ResponseContent:
		*self = append(*self, respContent.(ResponseContent))
	}
}

func (self *Queue) Pop() interface{} {
	if n := self.Len(); n == 0 {
		return nil
	} else {
		v := (*self)[n-1]
		*self = (*self)[:n-1]
		return v
	}
}

func (self *Queue) Tail() interface{} {
	return (*self)[self.Len()-1]
}

func (self *Queue) Len() int {
	return len(*self)
}

func (self *Queue) Self() interface{} {
	return *self
}
