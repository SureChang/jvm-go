package heap

import "classReader"

type Class struct {
	accessFlags uint16
	name string // thisClassName
	superClassName string
	interfaceNames []string
	constantPool *classReader.ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint
	staticSlotCount uint
	staticVars *Slots
}

func newClass(cf *classReader.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}