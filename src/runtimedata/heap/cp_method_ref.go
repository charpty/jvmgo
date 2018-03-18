package heap

import "classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(rtCp *ConstantPool, classInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = rtCp
	ref.className = classInfo.ClassName()
	ref.name, ref.descriptor = classInfo.NameAndDescriptor()
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolvedMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolvedMethodRef() {
	self.ResolvedClass()
	self.method = self.lookupMethod(self.class)
}

func (self *MethodRef) lookupMethod(cc *Class) *Method {
	for _, m := range cc.methods {
		if m.name == self.name && m.descriptor == self.descriptor {
			return m
		}
	}

	if sc := cc.superClass; sc != nil {
		return self.lookupMethod(sc)
	}
	return nil
}
