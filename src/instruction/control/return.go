package control

import (
	"instruction"
	"runtimedata"
)

type RETURN struct{ instruction.NoOperandsInstruction }

func (self *RETURN) Execute(frame *runtimedata.Frame) {
	frame.Thread().PopFrame()
}
