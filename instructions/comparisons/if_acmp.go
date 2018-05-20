package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//将栈顶的两个引用比较，符合条件就跳转指令
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, self.Offset)
	}
}
