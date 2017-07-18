package entry

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	return newDirEntry(path)
}
