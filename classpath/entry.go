package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) //路径分隔符，Windows是;  类Unix是:

//类路径项接口
type Entry interface {
	readClass(className string) ([]byte, Entry, error) //读取类数据
	String() string                                    //返回类路径名
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.Contains(path, "*") {
		return newWildCard(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
