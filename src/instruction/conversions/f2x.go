package conversions

import (
	"instruction"
	"runtimedata"
)
// Convert float to double
type F2D struct{ instruction.NoOperandsInstruction }

func (self *F2D) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct{ instruction.NoOperandsInstruction }

func (self *F2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct{ instruction.NoOperandsInstruction }

func (self *F2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
