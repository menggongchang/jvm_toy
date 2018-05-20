package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 | v2)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 | v2)
}
