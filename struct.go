package file

import "os"

// File 文件对象
type File struct {
	file   *os.File // 文件
	offset int64    // 偏移量
}
