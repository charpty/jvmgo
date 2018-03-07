package stores

import (
	"instruction"
	"runtimedata"
)

// Store double into local variable
type DSTORE struct{ instruction.Index8Instruction }

func (self *DSTORE) Execute(frame *runtimedata.Frame) {
	dstore(frame, uint(self.Index))
}

type DSTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_0) Execute(frame *runtimedata.Frame) {
	dstore(frame, 0)
}

type DSTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_1) Execute(frame *runtimedata.Frame) {
	dstore(frame, 1)
}

type DSTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_2) Execute(frame *runtimedata.Frame) {
	dstore(frame, 2)
}

type DSTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_3) Execute(frame *runtimedata.Frame) {
	dstore(frame, 3)
}

func dstore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
