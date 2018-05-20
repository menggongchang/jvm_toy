package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 ^ v2)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 ^ v2)
}
