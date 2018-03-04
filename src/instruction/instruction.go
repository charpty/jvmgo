package instruction

import "runtimedata"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame runtimedata.Frame)
}

type NoOperandsInstruction struct {
	// empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}
