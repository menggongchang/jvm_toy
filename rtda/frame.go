package rtda

//每个方法执行的同时会创建一个栈帧，
//栈帧用于存储局部变量表、操作数栈、动态链接、方法出口等信息。
//每个方法从调用直至执行完成的过程，就对应着一个栈帧在虚拟机栈中入栈到出栈的过程。
type Frame struct {
	lower        *Frame        //链表
	localVars    LocalVars     //局部变量表
	operandStack *OperandStack //操作数栈
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
