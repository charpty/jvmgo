package heap

import (
	"classfile"
)

type Method struct {
	ClassMember
	MaxLocals uint
	MaxStack  uint
	Code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	r := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		r[i] = &Method{}
		r[i].class = class
		r[i].copyMemberInfo(cfMethod)
		attr := cfMethod.CodeAttribute()
		if attr == nil {
			continue
		}
		r[i].MaxLocals = attr.MaxLocals()
		r[i].MaxStack = attr.MaxStack()
		r[i].Code = attr.Code()
	}
	return r
}
