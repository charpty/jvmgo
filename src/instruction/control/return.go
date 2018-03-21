package control

import (
	"instruction"
	"runtimedata"
)

type RETURN struct{ instruction.NoOperandsInstruction }
type ARETURN struct{ instruction.NoOperandsInstruction } // Return reference from method
type DRETURN struct{ instruction.NoOperandsInstruction } // Return double from method
type FRETURN struct{ instruction.NoOperandsInstruction } // Return float from method
type IRETURN struct{ instruction.NoOperandsInstruction } // Return int from method
type LRETURN struct{ instruction.NoOperandsInstruction }

func (self *RETURN) Execute(frame *runtimedata.Frame) {
	frame.Thread().PopFrame()
}

func (self *ARETURN) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	thread := frame.Thread()
	thread.PopFrame()
	thread.CurrentFrame().OperandStack().PushRef(ref)
}

func (self *DRETURN) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

func (self *FRETURN) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

func (self *IRETURN) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

func (self *LRETURN) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
