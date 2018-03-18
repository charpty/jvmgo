package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type NEW struct {
	instruction.Index16Instruction
}

func (self *NEW) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	cnt := cp.GetConstant(self.Index)
	cref := cnt.(*heap.ClassRef)
	class := cref.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("can not init interface or abstract class")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
