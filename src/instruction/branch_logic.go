package instruction

import "runtimedata"

func Branch(frame *runtimedata.Frame, offset int) {
	pc := frame.Thread().PC();
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
