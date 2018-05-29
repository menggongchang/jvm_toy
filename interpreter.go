package main

import (
	"fmt"
	"jvmgo/instructions"
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

//方法结束指令return还没有实现，必然报错
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//计算PC
		pc := frame.NextPC()
		thread.SetPC(pc)

		//解码指令
		reader.Reset(byteCode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//执行指令
		fmt.Printf("pc:%2d inst:%T %v \n", pc, inst, inst)
		inst.Execute(frame)
	}
}
