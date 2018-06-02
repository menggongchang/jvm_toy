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

//解析非接口方法的符号引用
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass() //先得到方法的类
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookUpMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookUpMethod(class *Class, name, descriptor string) *Method {
	method := LookUpMethodInclass(class, name, descriptor)
	if method == nil {
		method = LookUpMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
