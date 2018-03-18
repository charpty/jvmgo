package heap

import "classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(rtCp *ConstantPool, classInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = rtCp
	ref.className = classInfo.ClassName()
	ref.name, ref.descriptor = classInfo.NameAndDescriptor()
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolvedFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolvedFieldRef() {
	self.ResolvedClass()
	self.field = self.lookupField(self.class)
}

func (self *FieldRef) lookupField(cc *Class) *Field {
	for _, c := range cc.fields {
		if c.name == self.name && c.descriptor == self.descriptor {
			return c
		}
	}
	for _, intfc := range cc.interfaces {
		if r := self.lookupField(intfc); r != nil {
			return r
		}
	}
	if sc := cc.superClass; sc != nil {
		return self.lookupField(sc)
	}
	return nil

}
