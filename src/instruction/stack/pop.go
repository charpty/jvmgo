package stack

import (
	"instruction"
	"runtimedata"
)

type POP struct{ instruction.NoOperandsInstruction }
type POP2 struct{ instruction.NoOperandsInstruction }

func (self *POP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
