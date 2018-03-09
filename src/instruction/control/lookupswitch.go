package control

import (
	"instruction"
	"runtimedata"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *instruction.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *runtimedata.Frame) {
	key := frame.OperandStack().PopInt()
	var offset int
	for i, size := 0, len(self.matchOffsets); i <= size; i = i + 2 {
		if key == self.matchOffsets[i] {
			offset = int(self.matchOffsets[i+1])
			instruction.Branch(frame, offset)
			return
		}
	}
	instruction.Branch(frame, int(self.defaultOffset))
}
