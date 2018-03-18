package runtimedata

import "runtimedata/heap"

type Frame struct {
	// 栈中桢通过链表形式连接
	lower *Frame
	// 本地变量表
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
	// 所属线程
	thread *Thread
	// 当前帧所在方法
	method *heap.Method
	// 下一个执行指令位置
	nextPC int
}

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals),
		operandStack: newOperandStack(method.MaxStack),
	}
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
