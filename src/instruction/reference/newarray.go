package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

type NEW_ARRAY struct {
	atype uint8
}

type ARRAY_LENGTH struct {
	instruction.NoOperandsInstruction
}

func (self *NEW_ARRAY) FetchOperands(reader *instruction.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *runtimedata.Frame) {
	arrlen := frame.OperandStack().PopInt()
	if arrlen < 0 {
		panic("array length must greater than 0")
	}
	loader := frame.Method().Class().Loader()
	arrClass := getArrayClass(self.atype, loader)
	arr := arrClass.NewArray(uint(arrlen))
	frame.OperandStack().PushRef(arr)
}

func (self *ARRAY_LENGTH) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	arrLen := ref.ArrayLength()
	frame.OperandStack().PushInt(arrLen)
}

func getArrayClass(atype uint8, loader *heap.ClassLoader) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("unkonw array class type:" + string(atype))
	}
}
