package rtda

type Frame struct {
	lower *Frame
	thread *Thread
	localVars LocalVars
	operandStack *OperandStack
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread: thread,
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) Thread() *Thread {
	return self.thread
}


func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}