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
	if ref.className == "" {
		panic("ClassName can not be empty")
	}
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
	self.method = self.LookupMethodInClass(self.class)
	if self.method == nil {
		self.method = self.lookupMethodInInterfaces(self.class.interfaces, )
	}
}

func (self *MethodRef) LookupMethodInClass(cc *Class) *Method {
	for c := cc; c != nil; c = c.superClass {
		for _, m := range cc.methods {
			if m.name == self.name && m.descriptor == self.descriptor {
				return m
			}
		}
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
