package heap

import "classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
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
	self.method = self.lookupMethodInInterfaces(self.class.interfaces, )

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
