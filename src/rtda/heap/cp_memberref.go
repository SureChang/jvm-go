package heap

import "classReader"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classReader.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
