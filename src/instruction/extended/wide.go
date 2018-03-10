package extended

import (
	"instruction"
	"runtimedata"
	"instruction/loads"
	"instruction/stores"
	"instruction/math"
)

type WIDE struct {
	wideInstruction instruction.Instruction
}

func (self *WIDE) FetchOperands(reader *instruction.BytecodeReader) {
	tag := reader.ReadUint8()
	switch tag {
	case 0x5:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.wideInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.wideInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	default:
		panic("can not wide instruction: " + string(tag))
	}

}

func (self *WIDE) Execute(frame *runtimedata.Frame) {
	self.wideInstruction.Execute(frame)
}
