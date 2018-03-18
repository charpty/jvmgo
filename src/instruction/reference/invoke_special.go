package reference

import (
	"instruction"
	"runtimedata"
)

type INVOKE_SPECIAL struct{ instruction.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PopRef()
}
