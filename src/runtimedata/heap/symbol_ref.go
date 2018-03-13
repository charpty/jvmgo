package heap

type SymbolRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}
