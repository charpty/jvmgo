package comparisons

import (
	"instruction"
	"runtimedata"
)
// Compare double
type DCMPG struct{ instruction.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *runtimedata.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ instruction.NoOperandsInstruction }

func (self *DCMPL) Execute(frame *runtimedata.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *runtimedata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
