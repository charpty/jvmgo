package comparisons

import (
	"instruction"
	"runtimedata"
)

type IFEQ struct{ instruction.BranchInstruction }
type IFNE struct{ instruction.BranchInstruction }
type IFLT struct{ instruction.BranchInstruction }
type IFLE struct{ instruction.BranchInstruction }
type IFGT struct{ instruction.BranchInstruction }
type IFGE struct{ instruction.BranchInstruction }

func (self *IFEQ) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	if stack.PopInt() == 0 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IFNE) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IFLT) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IFLE) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IFGT) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IFGE) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		instruction.Branch(frame, self.Offset)
	}
}
