package math

import (
	"instructions/base"
	"rtda"
)

type IAND struct{ base.NoOperandsInstruction }

type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
