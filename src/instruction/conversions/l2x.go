package conversions

import (
	"instruction"
	"runtimedata"
)

// Convert long to double
type L2D struct{ instruction.NoOperandsInstruction }

func (self *L2D) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct{ instruction.NoOperandsInstruction }

func (self *L2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct{ instruction.NoOperandsInstruction }

func (self *L2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
