package instruction

import (
	"runtimedata"
	"runtimedata/heap"
	"util"
)

func InvokeMethod(invokerFrame *runtimedata.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	util.Debug("call method: " + method.Name())
	argSlotSlot := int(method.ArgCount)
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
