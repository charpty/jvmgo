package loads

import (
	"instruction"
	"runtimedata"
)


// Load double from local variable
type DLOAD struct{ instruction.Index8Instruction }

func (self *DLOAD) Execute(frame *runtimedata.Frame) {
	dload(frame, uint(self.Index))
}

type DLOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *runtimedata.Frame) {
	dload(frame, 0)
}

type DLOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *runtimedata.Frame) {
	dload(frame, 1)
}

type DLOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *runtimedata.Frame) {
	dload(frame, 2)
}

type DLOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *runtimedata.Frame) {
	dload(frame, 3)
}

func dload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
