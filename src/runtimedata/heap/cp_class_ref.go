package heap

import "classfile"

type ClassRef struct {
	SymbolRef
}

func newClassRef(rtCp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = rtCp
	ref.className = classInfo.Name()
	if ref.className == "" {
		panic("ClassName can not be empty")
	}
	return ref
}
