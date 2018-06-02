package rtda

type Stack struct {
	maxSize uint   //最多帧的个数
	size    uint   //当前帧的个数
	_top    *Frame //虚拟机栈通过链表实现
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverFlowError!")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = self._top.lower
	self.size--
	top.lower = nil
	return top
}
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}

func (self *Stack) IsEmpty() bool {
	return self._top == nil
}
