package constants

import (
	"instruction"
	"runtimedata"
)

// Push item from run-time constant pool
type LDC struct{ instruction.Index8Instruction }

func (self *LDC) Execute(frame *runtimedata.Frame) {
	_ldc(frame, self.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ instruction.Index16Instruction }

func (self *LDC_W) Execute(frame *runtimedata.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *runtimedata.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
		// case string:
		// case *heap.ClassRef:
		// case MethodType, MethodHandle
	default:
		
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ instruction.Index16Instruction }

func (self *LDC2_W) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
