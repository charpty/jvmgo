package constants

import (
	"instruction"
	"runtimedata"
)

// 都是一些常量指令，他们的功能就是将常量推入到操作数栈
// 这些指令本身没有操作数，指令下划线后就是其隐含的操作数，首字母代表类型
type ACONST_NULL struct{ instruction.NoOperandsInstruction }
type DCONST_0 struct{ instruction.NoOperandsInstruction }
type DCONST_1 struct{ instruction.NoOperandsInstruction }
type FCONST_0 struct{ instruction.NoOperandsInstruction }
type FCONST_1 struct{ instruction.NoOperandsInstruction }
type FCONST_2 struct{ instruction.NoOperandsInstruction }
type ICONST_M1 struct{ instruction.NoOperandsInstruction }
type ICONST_0 struct{ instruction.NoOperandsInstruction }
type ICONST_1 struct{ instruction.NoOperandsInstruction }
type ICONST_2 struct{ instruction.NoOperandsInstruction }
type ICONST_3 struct{ instruction.NoOperandsInstruction }
type ICONST_4 struct{ instruction.NoOperandsInstruction }
type ICONST_5 struct{ instruction.NoOperandsInstruction }
type LCONST_0 struct{ instruction.NoOperandsInstruction }
type LCONST_1 struct{ instruction.NoOperandsInstruction }

// ref null
func (self *ACONST_NULL) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushRef(nil)
}

// dobule 0
func (self *DCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (self *DCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// float 0
func (self *FCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

func (self *FCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

func (self *FCONST_2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

func (self *ICONST_M1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (self *ICONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(0)
}

// int 1
func (self *ICONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(1)
}

func (self *ICONST_2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(2)
}

func (self *ICONST_3) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(3)
}

func (self *ICONST_4) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(4)
}

func (self *ICONST_5) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(5)
}

// long 0
func (self *LCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(0)
}

func (self *LCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(1)
}
