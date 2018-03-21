package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type INVOKE_STATIC struct{ instruction.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	instruction.InvokeMethod(frame, resolvedMethod)
}
