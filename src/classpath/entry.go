package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	path = strings.ToLower(path);
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
