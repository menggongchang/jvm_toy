package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//取反
type INEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

type LNEG struct{ base.NoOperandsInstruction }

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

type FNEG struct{ base.NoOperandsInstruction }

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

type DNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}
