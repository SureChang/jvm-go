package heap

import "classReader"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classReader.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
