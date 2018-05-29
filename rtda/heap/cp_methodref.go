package heap

import (
	"jvmgo/classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, methodInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&methodInfo.ConstantMemberRefInfo)
	return ref
}
