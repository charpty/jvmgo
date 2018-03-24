package heap

import (
	"classfile"
)

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
	return ref
}

func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolvedInterfaceMethod()
	}
	return self.method
}

func (self *InterfaceMethodRef) resolvedInterfaceMethod() {
	self.ResolvedClass()
	self.method = self.lookupMethodInInterfaces([]*Class{self.class})
}

func (self *InterfaceMethodRef) lookupMethodInInterfaces(intfs []*Class) *Method {
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

func (self *InterfaceMethodRef) LookupMethodInClass(cc *Class) *Method {
	for _, m := range cc.methods {
		if m.name == self.name && m.descriptor == self.descriptor {
			return m
		}
	}
	if sc := cc.superClass; sc != nil {
		return self.LookupMethodInClass(sc)
	}
	return nil
}
