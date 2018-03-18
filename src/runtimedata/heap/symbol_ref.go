package heap

type SymbolRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymbolRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymbolRef) resolveClassRef() {
	cc := self.cp.class
	rc := cc.loader.LoadClass(self.className)
	if rc.isAccessibleTo(cc) {
		panic("java.lang.IllegalAccessError: " + rc.name + "is not accessible to" + cc.name)
	}
	self.class = rc
}
