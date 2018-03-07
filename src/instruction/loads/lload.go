package loads

import (
	"instruction"
	"runtimedata"
)

// Load long from local variable
type LLOAD struct{ instruction.Index8Instruction }

func (self *LLOAD) Execute(frame *runtimedata.Frame) {
	lload(frame, uint(self.Index))
}

type LLOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *runtimedata.Frame) {
	lload(frame, 0)
}

type LLOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *runtimedata.Frame) {
	lload(frame, 1)
}

type LLOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *runtimedata.Frame) {
	lload(frame, 2)
}

type LLOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *runtimedata.Frame) {
	lload(frame, 3)
}

func lload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
