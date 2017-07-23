package heap

import "classReader"

type MethodRef struct {
	MemberRef
	method *Method
}
func newMethodRef(cp *ConstantPool,
	refInfo *classReader.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
