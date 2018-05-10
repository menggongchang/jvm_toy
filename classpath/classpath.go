package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry //启动类路径
	extClassPath  Entry //扩展类路径
	userClassPath Entry //用户类路径
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (self *ClassPath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildCard(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildCard(jreExtPath)
}

func (self *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}

//获取jre目录
func getJreDir(jreOption string) string {
	//首先使用-Xjre选项作为jre目录
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	//如果没有,在当前目录寻找
	if exists("./jre") {
		return "./jre"
	}
	//如果找不到使用环境变量
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

//判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClassPath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}
