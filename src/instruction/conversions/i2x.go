package conversions

import (
	"instruction"
	"runtimedata"
)

// Convert int to byte
type I2B struct{ instruction.NoOperandsInstruction }

func (self *I2B) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// Convert int to char
type I2C struct{ instruction.NoOperandsInstruction }

func (self *I2C) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

// Convert int to short
type I2S struct{ instruction.NoOperandsInstruction }

func (self *I2S) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

// Convert int to long
type I2L struct{ instruction.NoOperandsInstruction }

func (self *I2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// Convert int to float
type I2F struct{ instruction.NoOperandsInstruction }

func (self *I2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// Convert int to double
type I2D struct{ instruction.NoOperandsInstruction }

func (self *I2D) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
