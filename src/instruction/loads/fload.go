package loads

import (
	"instruction"
	"runtimedata"
)


// Load float from local variable
type FLOAD struct{ instruction.Index8Instruction }

func (self *FLOAD) Execute(frame *runtimedata.Frame) {
	fload(frame, uint(self.Index))
}

type FLOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *FLOAD_0) Execute(frame *runtimedata.Frame) {
	fload(frame, 0)
}

type FLOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *FLOAD_1) Execute(frame *runtimedata.Frame) {
	fload(frame, 1)
}

type FLOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *FLOAD_2) Execute(frame *runtimedata.Frame) {
	fload(frame, 2)
}

type FLOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *FLOAD_3) Execute(frame *runtimedata.Frame) {
	fload(frame, 3)
}

func fload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
