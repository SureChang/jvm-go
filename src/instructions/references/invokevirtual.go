package references

import (
	"rtda"
	"instructions/base"
	"rtda/heap"
	"fmt"
	"strings"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ base.Index16Instruction }
// hack!
func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// hack System.out.println()
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V": fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V": fmt.Printf("%c\n", stack.PopInt())
	case "(B)V": fmt.Printf("%v\n", stack.PopInt())
	case "(S)V": fmt.Printf("%v\n", stack.PopInt())
	case "(I)V": fmt.Printf("%v\n", stack.PopInt())
	case "(F)V": fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V": fmt.Printf("%v\n", stack.PopLong())
	case "(D)V": fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		if(strings.Contains(goStr, "#gun")) {
			goStr = strings.Replace(goStr, "#gun", "▄︻┻═┳一", -1)
		}
		fmt.Println(goStr)
	default: panic("println: " + descriptor)
	}
	stack.PopRef()
}