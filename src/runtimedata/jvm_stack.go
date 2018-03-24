package runtimedata

type Stack struct {
	maxSize uint
	size    uint
	top     *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
		size:    0,
		top:     nil,
	}
}

func (self *Stack) Push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self.top != nil {
		frame.lower = self.top
	}
	self.top = frame
	self.size++
}

func (self *Stack) Pop() *Frame {
	if self.size == 0 {
		panic("thread stack is empty")
	}
	r := self.top
	self.top = r.lower
	r.lower = nil
	self.size--
	return r
}

func (self *Stack) Peek() *Frame {
	return self.top
}
