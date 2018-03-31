package main

import "flag"
import "fmt"
import (
	"os"
	"util"
	"strconv"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	// 是否开启debug日志
	debug      string
	cpOption   string
	XjreOption string
	class      string
	args       []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.debug, "debug", string(util.LOG_LEVEL_ERROR), "print debug log")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	if cmd.debug == "true" {
		util.LogLevel = util.LOG_LEVEL_DEBUG
	} else {
		level, err := strconv.Atoi(cmd.debug)
		if err == nil {
			util.LogLevel = level
		}
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
