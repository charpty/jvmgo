package loads

import (
	"instruction"
	"runtimedata"
)

// 加载命令，将指定索引的本地整型变量表变量加载到操作数栈中
type ILOAD struct{ instruction.Index8Instruction }

// 把推入操作数栈
type ILOAD_0 struct{ instruction.NoOperandsInstruction }
type ILOAD_1 struct{ instruction.NoOperandsInstruction }
type ILOAD_2 struct{ instruction.NoOperandsInstruction }
type ILOAD_3 struct{ instruction.NoOperandsInstruction }

func iload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *runtimedata.Frame) {
	iload(frame, self.Index)
}

func (self *ILOAD_0) Execute(frame *runtimedata.Frame) {
	iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *runtimedata.Frame) {
	iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *runtimedata.Frame) {
	iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *runtimedata.Frame) {
	iload(frame, 3)
}
