package control

import (
	"instruction"
	"runtimedata"
)

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.tableswitch
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *instruction.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	var offset int
	if val >= self.low && val <= self.high {
		offset = int(self.jumpOffsets[val-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	instruction.Branch(frame, offset)
}
