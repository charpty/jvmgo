package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type INVOKE_SPECIAL struct{ instruction.Index16Instruction }

// 调用私有方法和构造函数
// TODO 深究调用过程？
func (self *INVOKE_SPECIAL) Execute(frame *runtimedata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	resolvedClass := resolvedMethod.Class()

	if resolvedMethod.IsStatic() {
		panic("INVOKE_SPECIAL can not call static method")
	}

	invokedMethod := resolvedMethod

	if currentClass.IsSubClassOf(resolvedClass) && currentClass.IsSuper() && resolvedMethod.Name() != "<init>" {

	}

	instruction.InvokeMethod(frame, invokedMethod)

}
