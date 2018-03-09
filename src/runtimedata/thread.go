package runtimedata

// 一个线程持有唯一的PC计数器和线程栈
type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	var thread *Thread = &Thread{
		stack: NewStack(1024),
	}
	return thread
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.Push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.Pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.Peek()
}

func (self *Thread) NewFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		thread:       self,
	}
}
