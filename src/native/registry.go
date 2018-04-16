package native

import (
	"runtimedata"
)

type NativeMethod func(frame *runtimedata.Frame)

var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *runtimedata.Frame) {
	// do nothing
}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := getMethodSignature(className, methodName, methodDescriptor)
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := getMethodSignature(className, methodName, methodDescriptor)
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func getMethodSignature(className, methodName, methodDescriptor string) string {
	return className + "~" + methodName + "~" + methodDescriptor
}