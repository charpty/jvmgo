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

func (self *BIPUSH) FetchOperand(reader *instruction.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *runtimedata.Frame) {
	r := int32(self.val)
	frame.OperandStack().PushInt(r)
}
