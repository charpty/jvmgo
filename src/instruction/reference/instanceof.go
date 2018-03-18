package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

// Determine if object is of given type
type INSTANCE_OF struct{ instruction.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
