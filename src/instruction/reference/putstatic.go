package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type PUT_STATIC struct {
	instruction.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *runtimedata.Frame) {
	method := frame.Method()
	cp := method.Class().ConstantPool()
	field := cp.GetConstant(self.Index).(*heap.FieldRef).ResolvedField()
	descriptor := field.Descriptor()

	class := field.Class()
	stack := frame.OperandStack()
	slots := class.StaticVars()
	slotId := field.SlotId()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
