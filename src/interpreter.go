package main

import (
	"runtimedata"
	"instruction"
	"instruction/factory"
	"util"
	"runtimedata/heap"
)

func interpret(method *heap.Method) {
	thread := runtimedata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread)
}

func loop(thread *runtimedata.Thread) {
	reader := &instruction.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		util.Debug("run current frame: " + frame.Method().Name())
		bytecode := frame.Method().Code
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := factory.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		util.Debug("pc:%2d inst:%T %v", pc, inst, inst)
		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}

func catchErr(frame *runtimedata.Frame) {
	if r := recover(); r != nil {
		util.Error("LocalVars:%v", frame.LocalVars())
		util.Error("OperandStack:%v", frame.OperandStack())
		panic(r)
	}
}
