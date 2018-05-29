package heap

import (
	"fmt"
	"jvmgo/classfile"
)

type Constant interface {
}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfcp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfcp)
	consts := make([]Constant, cpCount)
	constantPool := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfcp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()

		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(constantPool, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(constantPool, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(constantPool, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(constantPool, methodrefInfo)
		default:
			// todo
		}
	}
	return constantPool
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	c := self.consts[index]
	if c != nil {
		return c
	}

	panic(fmt.Sprintf("No constants at index %d", index))
}
