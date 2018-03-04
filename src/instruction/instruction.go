package instruction

import "runtimedata"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame runtimedata.Frame)
}
