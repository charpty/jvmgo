package heap

import "classfile"

type Method struct {
	ClassMember
	MaxLocals uint
	MaxStack  uint
	code      []byte
}

func newMethod(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	r := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		r[i] = &Method{}
		r[i].class = class
		r[i].copyMemberInfo(cfMethod)
	}
	return r
}
