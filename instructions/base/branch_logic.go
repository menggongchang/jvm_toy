package base

import (
	"jvmgo/rtda"
)

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPc := pc + offset
	frame.SetNextPC(nextPc) //5.12
}
