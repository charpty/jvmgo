package heap

import "classfile"

type FieldRef struct {
	MemberRef
}

func newFieldRef(rtCp *ConstantPool, classInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = rtCp
	ref.className = classInfo.ClassName()
	ref.name, ref.descriptor = classInfo.NameAndDescriptor()
	return ref
}
