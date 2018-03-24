package runtimedata

import "math"
import "runtimedata/heap"

type OperandStack struct {
	size  uint
	slots []heap.Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]heap.Slot, maxStack),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32) {
	slot := &self.slots[self.size]
	slot.SetNum(val)
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].Num()
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	(&self.slots[self.size]).SetNum(int32(bits))
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].Num())
	return math.Float32frombits(bits)
}

// long consumes two slots
func (self *OperandStack) PushLong(val int64) {
	(&self.slots[self.size]).SetNum(int32(val))
	(&self.slots[self.size+1]).SetNum(int32(val >> 32))
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].Num())
	high := uint32(self.slots[self.size+1].Num())
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].SetRef(ref)
	self.size++
}

func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].Ref()
	self.slots[self.size].SetRef(nil)
	return ref
}

func (self *OperandStack) PopSlot() heap.Slot {
	self.size--;
	return self.slots[self.size];
}

func (self *OperandStack) PushSlot(slot heap.Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	u := self.size - 1 - n
	if u >= self.size {
		return nil
	}
	return self.slots[u].Ref()
}
