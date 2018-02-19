package classpath

import "os"
import "strings"

// 这是系统路径分隔符，在unix下是冒号，在windows下是分号
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
    // go语言里接口不用显式实现，只要方法匹配即可
	readClass(className string) ([]byte, Entry, error)
    // 类似于java的toString()方法, 方法名首字母大写代表公开访问
	String() string
}

// 获取类字节流都会调用此方法, 只能在包内使用
func newEntry(path string) Entry {
    // 如果含有分隔符就代表是多个classpath
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

    // 通用匹配
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

    // jar包
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

    // 普通目录
	return newDirEntry(path)
}
