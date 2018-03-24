package heap

type MemberRef struct {
	SymbolRef
	name       string
	descriptor string
}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
