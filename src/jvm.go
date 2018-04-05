package main

import "classpath"
import "runtimedata/heap"

import (
	"fmt"
	"util"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	util.Debug("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)
	classloader := heap.NewClassLoader(cp)
	class := classloader.LoadClass(cmd.class)

	mainMethod := class.GetMainMethod()
	interpret(mainMethod, cmd.args)
}
