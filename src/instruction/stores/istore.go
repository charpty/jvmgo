package stores

import (
	"instruction"
	"runtimedata"
)

// Store int into local variable
type ISTORE struct{ instruction.Index8Instruction }

func (self *ISTORE) Execute(frame *runtimedata.Frame) {
	istore(frame, uint(self.Index))
}

type ISTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_0) Execute(frame *runtimedata.Frame) {
	istore(frame, 0)
}

type ISTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_1) Execute(frame *runtimedata.Frame) {
	istore(frame, 1)
}

type ISTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_2) Execute(frame *runtimedata.Frame) {
	istore(frame, 2)
}

type ISTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_3) Execute(frame *runtimedata.Frame) {
	istore(frame, 3)
}

func istore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
