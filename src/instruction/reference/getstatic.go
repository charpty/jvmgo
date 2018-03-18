package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type GET_STATIC struct {
	instruction.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *runtimedata.Frame) {
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
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
