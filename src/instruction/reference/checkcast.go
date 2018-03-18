package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

// Check whether object is of given type
type CHECK_CAST struct{ instruction.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
