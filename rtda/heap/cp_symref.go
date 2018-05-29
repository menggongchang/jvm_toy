package heap

type SymRef struct {
	cp        *ConstantPool
	class     *Class
	className string
}

//类符号引用解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.ResolvedClassRef()
	}
	return self.class
}

func (self *SymRef) ResolvedClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
