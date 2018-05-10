package main

import (
	"fmt"
)

//Usage: %s [-options] class [args...]
func main() {
	cmd := parseCmd()
	if cmd.helpFlag {
		printUsage()
	} else if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s, class: %s, args: %v \n", cmd.cpOption, cmd.class, cmd.args)
}
