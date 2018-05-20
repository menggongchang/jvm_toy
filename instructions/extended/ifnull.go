package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//if ref is null branch
type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	if frame.OperandStack().PopRef() == nil {
		base.Branch(frame, self.Offset)
	}
}

//if ref is not null branch
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	if frame.OperandStack().PopRef() != nil {
		base.Branch(frame, self.Offset)
	}
}
