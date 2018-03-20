package heap

import (
	"classfile"
)

type Method struct {
	ClassMember
	MaxLocals uint
	MaxStack  uint
	Code      []byte
	ArgCount  uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	r := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		r[i] = &Method{}
		r[i].class = class
		r[i].copyMemberInfo(cfMethod)
		r[i].calcArgCount(cfMethod)
		attr := cfMethod.CodeAttribute()
		if attr != nil {
			r[i].MaxLocals = attr.MaxLocals()
			r[i].MaxStack = attr.MaxStack()
			r[i].Code = attr.Code()
		}
	}
	return r
}
func (self *Method) calcArgCount(cfMethod *classfile.MemberInfo) {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.ArgCount++
		if paramType == "J" || paramType == "D" {
			self.ArgCount++
		}
	}
	if !self.IsStatic() {
		self.ArgCount++ // `this` reference
	}
}
