package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type MULTI_A_NEW_ARRAY struct {
	// 数组类型class常量池索引
	index     uint16
	dimension uint8
}

func (self *MULTI_A_NEW_ARRAY) FetchOperands(reader *instruction.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimension = reader.ReadUint8()
}
func (self *MULTI_A_NEW_ARRAY) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := self.popDimensions(stack)
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func newMultiDimensionalArray(counts []int32, class *heap.Class) *heap.Object {
	first := uint(counts[0])
	arr := class.NewArray(first)
	if len(counts) == 1 {
		return arr
	}
	refs := arr.Refs()
	for i := range refs {
		refs[i] = newMultiDimensionalArray(counts[1:], class)
	}
	return arr
}

func (self *MULTI_A_NEW_ARRAY) popDimensions(stack *runtimedata.OperandStack) []int32 {
	result := make([]int32, self.dimension)
	for i := 0; i < int(self.dimension); i++ {
		c := stack.PopInt()
		if c < 0 {
			panic("arr length must >= 0")
		}
		result[i] = c
	}
	return result
}
