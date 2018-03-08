package math

import (
	"instruction"
	"runtimedata"
)

// Boolean AND int
type IAND struct{ instruction.NoOperandsInstruction }

func (self *IAND) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ instruction.NoOperandsInstruction }

func (self *LAND) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
