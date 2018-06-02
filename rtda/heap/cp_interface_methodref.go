package heap

import (
	"jvmgo/classfile"
)

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

//解析接口方法的符号引用
func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass() //先得到方法的接口
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookUpInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookUpInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.Name() == name && method.Descriptor() == descriptor {
			return method
		}
	}
	method := LookUpMethodInInterfaces(iface.interfaces, name, descriptor)
	return method
}
