package math

import (
	"instruction"
	"runtimedata"
)
// Divide double
type DDIV struct{ instruction.NoOperandsInstruction }

func (self *DDIV) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type FDIV struct{ instruction.NoOperandsInstruction }

func (self *FDIV) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type IDIV struct{ instruction.NoOperandsInstruction }

func (self *IDIV) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// Divide long
type LDIV struct{ instruction.NoOperandsInstruction }

func (self *LDIV) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
