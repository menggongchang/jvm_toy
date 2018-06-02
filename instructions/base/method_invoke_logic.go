package base

import (
	"fmt"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	//传递方法参数
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
