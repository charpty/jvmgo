package main

import (
	"runtimedata"
	"fmt"
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
		util.Debug("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}

func catchErr(frame *runtimedata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
