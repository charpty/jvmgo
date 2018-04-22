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
		r[i] = newMethod(class, cfMethod)
	}
	return r
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	result := &Method{}
	result.class = class
	result.copyMemberInfo(cfMethod)
	result.copyAttributes(cfMethod)
	md := parseMethodDescriptor(result.descriptor)
	result.calcArgCount(cfMethod)
	attr := cfMethod.CodeAttribute()
	if attr != nil {
		result.MaxLocals = attr.MaxLocals()
		result.MaxStack = attr.MaxStack()
		result.Code = attr.Code()
	}
	if result.IsNative() {
		result.injectCodeAttribute(md.returnType)
	}
	return result
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.MaxStack = 4 // todo
	// 本地方法只需要有返回就可以
	self.MaxLocals = self.ArgCount
	switch returnType[0] {
	case 'V':
		self.Code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.Code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.Code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.Code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.Code = []byte{0xfe, 0xad} // lreturn
	default:
		self.Code = []byte{0xfe, 0xac} // ireturn
	}
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.MaxStack = codeAttr.MaxStack()
		self.MaxLocals = codeAttr.MaxLocals()
		self.Code = codeAttr.Code()
	}
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

func (self *Method) Signature() string {
	return self.class.name + "#" + self.name + self.descriptor
}
