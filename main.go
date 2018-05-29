package main

import (
	"fmt"
	"jvmgo/classpath"
	"jvmgo/rtda/heap"
	"strings"
)

//Usage: %s [-options] class [args...]
func main() {
	cmd := parseCmd()
	if cmd.helpFlag {
		printUsage()
	} else if cmd.versionFlag {
		fmt.Println("version 0.0.6")
	} else if cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %v, class: %s, args: %v \n", cp, cmd.class, cmd.args)

	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)

	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod) //执行main方法
	} else {
		fmt.Printf("Main method not found in class %s\n", className)
	}
}
