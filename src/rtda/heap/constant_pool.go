package heap

import (
	"fmt"
	"classReader"
)

type Constant interface{}

type ConstantPool struct {
	class *Class
	consts []Constant
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}

func newConstantPool(class *Class, cfCp classReader.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classReader.ConstantIntegerInfo:
			intInfo := cpInfo.(*classReader.ConstantIntegerInfo)
			consts[i] = intInfo.Value() // int32
		case *classReader.ConstantFloatInfo:
			floatInfo := cpInfo.(*classReader.ConstantFloatInfo)
			consts[i] = floatInfo.Value() // float32
		case *classReader.ConstantLongInfo:
			longInfo := cpInfo.(*classReader.ConstantLongInfo)
			consts[i] = longInfo.Value() // int64
			i++
		case *classReader.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classReader.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() // float64
			i++
		case *classReader.ConstantStringInfo:
			stringInfo := cpInfo.(*classReader.ConstantStringInfo)
			consts[i] = stringInfo.String() // string
		case *classReader.ConstantClassInfo:
			classInfo := cpInfo.(*classReader.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classReader.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classReader.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classReader.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classReader.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classReader.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classReader.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		}
	}
	return rtCp
}