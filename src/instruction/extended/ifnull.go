package extended

import (
	"instruction"
	"runtimedata"
)

// Branch if reference is null
type IFNULL struct{ instruction.BranchInstruction }

func (self *IFNULL) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		instruction.Branch(frame, self.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ instruction.BranchInstruction }

func (self *IFNONNULL) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		instruction.Branch(frame, self.Offset)
	}
}
