package instruction

type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint8() uint8 {
	r := self.code[self.pc]
	self.pc++
	return r
}
