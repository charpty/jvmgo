package heap

import "classfile"

type Field struct {
	ClassMember
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	r := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		r[i] = &Field{}
		r[i].class = class
		r[i].copyMemberInfo(cfField)
	}
	return r
}
