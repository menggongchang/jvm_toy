package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//当两个变量至少有一个是NAN，无法比较时，返回1或-1
type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	d2 := stack.PopDouble()
	d1 := stack.PopDouble()
	if d1 > d2 {
		stack.PushInt(1)
	} else if d1 == d2 {
		stack.PushInt(0)
	} else if d1 < d2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
