package main

import (
	"flag" //命令行解析包
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool     //输出帮助信息，然后退出
	versionFlag bool     //输出版本信息，然后退出
	cpOption    string   //指定用户类路径
	XjreOption  string   //指定jre目录，加载类
	class       string   //主类名
	args        []string //传递给主类的参数
}

//解析命令行
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage //解析失败时，打印用法
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
