package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) SetRefValue(fieldName string, fieldDescriptor string, value *Object) {
	field := self.class.GetField(fieldName, fieldDescriptor)
	self.data.(Slots).SetRef(field.SlotId(), value)
}

func (self *Object) GetRefValue(fieldName string, fieldDescriptor string) *Object {
	field := self.class.GetField(fieldName, fieldDescriptor)
	return self.data.(Slots).GetRef(field.SlotId())
}

func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}
