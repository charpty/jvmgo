package constants

import (
	"instruction"
	"runtimedata"
)

// 第一字符表示操作数的类型
// 这个指令是从操作数栈中获取一个byte类型，并转换为int再压入操作数栈顶
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *instruction.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *runtimedata.Frame) {
	r := int32(self.val)
	frame.OperandStack().PushInt(r)
}

type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *instruction.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *runtimedata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
