package reserved

import (
	"instruction"
	"native"
	"runtimedata"
)

type INVOKE_NATIVE struct{ instruction.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *runtimedata.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
