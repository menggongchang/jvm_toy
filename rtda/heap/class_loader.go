package heap

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
)

type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class //相当于方法区
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	cl := &ClassLoader{}
	cl.cp = cp
	cl.classMap = make(map[string]*Class)
	return cl
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name) //读取class文件，加载数据到内存
	class := self.defineClass(data)     //解析class文件，生成类数据，放入方法区
	link(class)                         //链接
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(className string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(className)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + className)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(classData []byte) *Class {
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	// printClassInfo(cf)
	return newClass(cf)
}

//测试
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	class.interfaces = make([]*Class, interfaceCount)
	for i, interfaceName := range class.interfaceNames {
		class.interfaces[i] = class.loader.LoadClass(interfaceName)
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//do nothing
}

//给类变量分配空间并给与初始值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class) //实例字段
	calcStaticFieldSlotIds(class)   //静态字段
	allocAndInitStaticVars(class)   //给类变量分配空间
}

func calcInstanceFieldSlotIds(class *Class) {
	slotID := uint(0)
	if class.superClass != nil {
		slotID = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotID = slotID
			slotID++
			if field.isLongOrDouble() {
				slotID++
			}
		}
	}
	class.instanceSlotCount = slotID
}

func calcStaticFieldSlotIds(class *Class) {
	slotID := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotID = slotID
			slotID++
			if field.isLongOrDouble() {
				slotID++
			}
		}
	}
	class.staticSlotCount = slotID
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotID := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotID, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotID, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotID, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotID, val)
		case "Ljava/lang/String;":
			panic("todo") //第八章实现
		}
	}
}
