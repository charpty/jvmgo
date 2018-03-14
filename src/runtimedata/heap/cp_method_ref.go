package heap

import "classfile"

type MethodRef struct {
	MemberRef
}

func newMethodRef(rtCp *ConstantPool, classInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = rtCp
	ref.className = classInfo.ClassName()
	ref.name, ref.descriptor = classInfo.NameAndDescriptor()
	return ref
}