package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//交换栈顶两个元素
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
