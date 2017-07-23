package heap

import "classReader"

type ClassMember struct {
	accessFlags uint16
	name string
	descriptor string
	class *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classReader.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}