package heap

import "classfile"

type Field struct {
	ClassMember
	constantValueIndex uint
	slotId             uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	r := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		r[i] = &Field{}
		r[i].class = class
		r[i].copyMemberInfo(cfField)
		// 从常量属性表中取值
		if attr := cfFields[i].ConstantAttribute(); attr != nil {
			r[i].constantValueIndex = uint(attr.ConstantValueIndex())
		}
	}
	return r
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "D" || self.descriptor == "J"
}

func (self *Field) ConstValueIndex() uint {
	return self.constantValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}
