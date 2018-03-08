package math

import (
	"instruction"
	"runtimedata"
)
// Negate double
type DNEG struct{ instruction.NoOperandsInstruction }

func (self *DNEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ instruction.NoOperandsInstruction }

func (self *FNEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ instruction.NoOperandsInstruction }

func (self *INEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ instruction.NoOperandsInstruction }

func (self *LNEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
