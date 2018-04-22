package native

import (
	"runtimedata"
	"util"
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
	util.Error("native method don't have: %s", key)
	util.Debug("current native map: %v", registry)

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func getMethodSignature(className, methodName, methodDescriptor string) string {
	return className + "~" + methodName + "~" + methodDescriptor
}
