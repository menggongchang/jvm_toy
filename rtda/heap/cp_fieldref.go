package heap

import (
	"jvmgo/classfile"
)

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, fieldInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	fieldRef := &FieldRef{}
	fieldRef.cp = cp
	fieldRef.copyMemberRefInfo(&fieldInfo.ConstantMemberRefInfo)
	return fieldRef
}

//解析字段的符号引用
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.ResolvedFieldRef()
	}
	return self.field
}

func (self *FieldRef) ResolvedFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookUpField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

//查找字段
func lookUpField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookUpField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookUpField(c.superClass, name, descriptor)
	}
	return nil
}
