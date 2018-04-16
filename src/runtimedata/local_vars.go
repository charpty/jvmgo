package runtimedata

import "math"
import "runtimedata/heap"

type LocalVars []heap.Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]heap.Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].SetNum(val)
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].Num()
}

func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].SetNum(int32(bits))
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].Num())
	return math.Float32frombits(bits)
}

// long consumes two slots
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].SetNum(int32(val))
	self[index+1].SetNum(int32(val >> 32))
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].Num())
	high := uint32(self[index+1].Num())
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].SetRef(ref)
}

func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].Ref()
}

func (self LocalVars) SetSlot(index uint, slot heap.Slot) {
	self[index] = slot
}

func (self LocalVars) GetSlot(index uint) heap.Slot {
	return self[index]
}

func (self LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}
