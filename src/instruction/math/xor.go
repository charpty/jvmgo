package math

import (
	"instruction"
	"runtimedata"
)

// Boolean XOR int
type IXOR struct{ instruction.NoOperandsInstruction }

func (self *IXOR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct{ instruction.NoOperandsInstruction }

func (self *LXOR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
