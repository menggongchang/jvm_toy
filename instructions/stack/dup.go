package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//复制栈顶变量
type DUP struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	slotc := stack.PopSlot()
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
	stack.PushSlot(slotc)
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	slotc := stack.PopSlot()
	slotd := stack.PopSlot()
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
	stack.PushSlot(slotd)
	stack.PushSlot(slotc)
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
}
