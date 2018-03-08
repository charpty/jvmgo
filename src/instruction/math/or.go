package math

import (
	"instruction"
	"runtimedata"
)
// Boolean OR int
type IOR struct{ instruction.NoOperandsInstruction }

func (self *IOR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ instruction.NoOperandsInstruction }

func (self *LOR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
