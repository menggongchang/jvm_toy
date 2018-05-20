package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//弹出操作数栈，一个位置
type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

//弹出操作数栈，2个位置例如double，long
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
