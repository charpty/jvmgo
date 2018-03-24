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

	thisObject := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgCount - 1)
	if thisObject == nil {
		panic("java.lang.NullPointerException")
	}

	// 如果方法是属于父类的方法，而实际调用者是子类，那么需要查找实际的方法
	// 因为有可能当前类中也有这个方法，那么就要调用子类的
	// 注意这里还是静态的方法查找模式，区别与运行时动态绑定的INVOKE_VIRTUAL
	if currentClass.IsSubClassOf(resolvedClass) && currentClass.IsSuper() && resolvedMethod.Name() != "<init>" {
		invokedMethod = methodRef.LookupMethodInClass(currentClass)
	}
	instruction.InvokeMethod(frame, invokedMethod)
}
