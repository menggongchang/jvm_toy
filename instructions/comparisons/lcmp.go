package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//compare long
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l2 := stack.PopLong()
	l1 := stack.PopLong()
	if l1 > l2 {
		stack.PushInt(1)
	} else if l1 == l2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
