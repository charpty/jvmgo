package reference

import (
	"instruction"
	"runtimedata"
	"runtimedata/heap"
)

type GET_FIELD struct {
	instruction.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *runtimedata.Frame) {
	// 实例引用是从栈上弹出的
	// 第一个操作数表明Field描述
	cp := frame.Method().Class().ConstantPool()
	field := cp.GetConstant(self.Index).(*heap.FieldRef).ResolvedField()

	slotId := field.SlotId()
	descriptor := field.Descriptor()
	stack := frame.OperandStack()
	instance := stack.PopRef()
	if instance == nil {
		panic("Null Exception")
	}
	slots := instance.Fields()
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
