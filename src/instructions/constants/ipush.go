package constants

import (
	"instructions/base"
	"rtda"
)

type BIPUSH struct { val int8 } // Push byte
type SIPUSH struct { val int16 } // Push short

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
