package main

import "classpath"
import "classfile"

import (
	"fmt"
	"strings"
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
	classData := readClassData(cmd)
	parseClassData(classData)
}

func parseClassData(classData []byte) {
	classFile, err := classfile.Parse(classData)
	if err != nil {
		fmt.Printf("Parse class error %s\n", err.Error())
		panic("Parse class error")
	}
	fmt.Printf("class file: %+v", classFile)
}

func readClassData(cmd *Cmd) (classData []byte) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		panic("Load class error, can not find the class file")
	}
	fmt.Printf("class data:%v\n", classData)
	return classData
}
