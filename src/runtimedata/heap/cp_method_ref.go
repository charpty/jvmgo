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
	if self.method == nil {
		self.method = self.lookupMethodInInterfaces(self.class.interfaces, )
	}
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

func (self *MethodRef) lookupMethodInInterfaces(intfs []*Class) *Method {
	for _, i := range intfs {
		for _, method := range i.methods {
			if method.name == self.name && method.descriptor == self.descriptor {
				return method
			}
		}
		method := self.lookupMethodInInterfaces(i.interfaces)
		if method != nil {
			return method
		}
	}

	return nil
}

func (self *MethodRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
