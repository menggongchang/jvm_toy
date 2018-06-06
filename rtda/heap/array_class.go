package heap

//数组类的特殊方法
func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("Nor array class：" + self.Name())
	}
	switch self.Name() {
	case "[Z": //byte
		return &Object{self, make([]int8, count)}
	case "[S": //short
		return &Object{self, make([]int16, count)}
	case "[I": //int
		return &Object{self, make([]int32, count)}
	case "[J": //long
		return &Object{self, make([]int64, count)}
	case "[B": //boolean
		return &Object{self, make([]int8, count)}
	case "[C": //char
		return &Object{self, make([]uint16, count)}
	case "[F": //float
		return &Object{self, make([]float32, count)}
	case "[D": //double
		return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}
