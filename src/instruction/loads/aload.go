package loads

import (
	"instruction"
	"runtimedata"
)

// Load reference from local variable
type ALOAD struct{ instruction.Index8Instruction }

func (self *ALOAD) Execute(frame *runtimedata.Frame) {
	aload(frame, uint(self.Index))
}

type ALOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *runtimedata.Frame) {
	aload(frame, 0)
}

type ALOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *runtimedata.Frame) {
	aload(frame, 1)
}

type ALOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *runtimedata.Frame) {
	aload(frame, 2)
}

type ALOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *runtimedata.Frame) {
	aload(frame, 3)
}

func aload(frame *runtimedata.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
