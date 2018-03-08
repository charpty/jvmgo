package stack

import (
	"instruction"
	"runtimedata"
)

// Duplicate the top operand stack value
type DUP struct{ instruction.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
             \_
               |
               V
[...][c][b][a][a]
*/
func (self *DUP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ instruction.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
func (self *DUP_X1) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
