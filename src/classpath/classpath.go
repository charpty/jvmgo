package classpath

import "os"
import "path/filepath"
import "util"

// 一共3种类路径
type Classpath struct {
	// 启动类路径，在HotSpot虚拟机实现中由Bootstrap ClassLoader加载
	bootClasspath Entry
	// 扩展类路径，由Extention ClassLoader加载
	extClasspath Entry
	// 用户类路径或者系统类路径，由AppClassLoader加载
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// 解析JDK核心类库，HotSpot使用C++实现加载rt.jar,i18n.jar等核心包
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	// 使用模糊匹配形式entry
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// 加载JDK扩展类库
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	javaHome := os.Getenv("JAVA_HOME")
	util.Debug("JAVA_HOME=" + javaHome)
	if jh := javaHome; jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
