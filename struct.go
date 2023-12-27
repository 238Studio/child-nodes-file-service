package file

import "os"

// File 文件对象
type File struct {
	File *os.File // 文件
}
