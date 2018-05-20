package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//compare long
//当两个变量至少有一个是NAN，无法比较时，返回1或-1
type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	f2 := stack.PopFloat()
	f1 := stack.PopFloat()
	if f1 > f2 {
		stack.PushInt(1)
	} else if f1 == f2 {
		stack.PushInt(0)
	} else if f1 < f2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
