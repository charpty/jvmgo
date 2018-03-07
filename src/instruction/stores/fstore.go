package stores

import (
	"instruction"
	"runtimedata"
)

// Store float into local variable
type FSTORE struct{ instruction.Index8Instruction }

func (self *FSTORE) Execute(frame *runtimedata.Frame) {
	fstore(frame, uint(self.Index))
}

type FSTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_0) Execute(frame *runtimedata.Frame) {
	fstore(frame, 0)
}

type FSTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_1) Execute(frame *runtimedata.Frame) {
	fstore(frame, 1)
}

type FSTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_2) Execute(frame *runtimedata.Frame) {
	fstore(frame, 2)
}

type FSTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_3) Execute(frame *runtimedata.Frame) {
	fstore(frame, 3)
}

func fstore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
