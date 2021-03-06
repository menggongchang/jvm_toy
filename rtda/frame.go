package rtda

import (
	"jvmgo/rtda/heap"
)

//每个方法执行的同时会创建一个栈帧，
//栈帧用于存储局部变量表、操作数栈、动态链接、方法出口等信息。
//每个方法从调用直至执行完成的过程，就对应着一个栈帧在虚拟机栈中入栈到出栈的过程。
type Frame struct {
	lower        *Frame        //链表
	localVars    LocalVars     //局部变量表
	operandStack *OperandStack //操作数栈
	thread       *Thread       //用于实现指令跳转
	method       *heap.Method
	nextPC       int //用于实现指令跳转
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
