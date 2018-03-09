package conversions

import (
	"instruction"
	"runtimedata"
)

// Convert double to float
type D2F struct{ instruction.NoOperandsInstruction }

func (self *D2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ instruction.NoOperandsInstruction }

func (self *D2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ instruction.NoOperandsInstruction }

func (self *D2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
