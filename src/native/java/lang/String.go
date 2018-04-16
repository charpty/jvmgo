package lang

import (
	"native"
	"runtimedata"
	"runtimedata/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *runtimedata.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
