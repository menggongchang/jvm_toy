package base

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

//执行类的初始化方法
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetInitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

//执行超类的初始化
func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
