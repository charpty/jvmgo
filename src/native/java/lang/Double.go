package lang

import (
	"math"
	"native"
	"runtimedata"
)

// 基于某一个类的native方法
const jlDouble = "java/lang/Double"

func init() {
	native.Register(jlDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register(jlDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *runtimedata.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value) // todo
	frame.OperandStack().PushLong(int64(bits))
}

func longBitsToDouble(frame *runtimedata.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits)) // todo
	frame.OperandStack().PushDouble(value)
}
