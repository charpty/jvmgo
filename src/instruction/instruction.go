package instruction

import "runtimedata"

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *runtimedata.Frame)
}

// 抽象无操作数类指令
type NoOperandsInstruction struct {
	// empty
}

// 抽象跳转类指令
type BranchInstruction struct {
	Offset int
}

// 抽象存储和加载类指令
type Index8Instruction struct {
	// 这些指令都需要存取局部变量表
	Index uint
}

// 抽象需要访问常量池指令
type Index16Instruction struct {
	Index uint
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
