package reference

import (
	"fmt"
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ instruction.Index16Instruction }

// hack!
func (self *INVOKE_VIRTUAL) Execute(frame *runtimedata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgCount - 1)
	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}

		panic("java.lang.NullPointerException")
	}
	invokedMethod := methodRef.LookupMethodInClass(ref.Class())

	if invokedMethod.IsAbstract() {
		panic("Can not call abstract method")
	}
	instruction.InvokeMethod(frame, invokedMethod)
}

// hack!
func _println(stack *runtimedata.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		fmt.Println(goStr)
	default:
		panic("println: %s" + descriptor)
	}
	stack.PopRef()
}
