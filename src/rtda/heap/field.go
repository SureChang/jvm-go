package heap

import "classReader"

type Field struct {
	ClassMember
}

func newFields(class *Class, cfFields []*classReader.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}
