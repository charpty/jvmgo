package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type ANEW_ARRAY struct {
	instruction.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *runtimedata.Frame) {
	arrLen := frame.OperandStack().PopInt()
	cp := frame.Method().Class().ConstantPool()
	arrClassRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	arrClass := arrClassRef.ResolvedClass()
	arr := arrClass.NewArray(uint(arrLen))
	frame.OperandStack().PushRef(arr)
}
