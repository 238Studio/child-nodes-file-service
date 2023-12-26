package file

// GetFile 获取文件
// 传入参数：文件路径
// 返回参数：os.File, error
func GetFile(path string) (File, error) {
	file, err := getFile(path)
	//获取文件偏移量
	offset, err := file.Seek(0, 2)
	if err != nil {
		return File{}, err
	}
	return File{file: file, offset: offset}, err
}

// CreateFile 创建文件
// 传入参数：文件路径
// 返回参数：os.File, error
func CreateFile(path string) (File, error) {
	file, err := createFile(path)
	return File{file: file, offset: 0}, err
}

// WriteFile 写入文件
// 传入参数：文件路径、文件内容([]byte)
// 返回参数：错误信息
func (file *File) WriteFile(content []byte) error {
	err := writeFileByOffset(file.file, content, 0)
	if err != nil {
		return err
	}

	file.offset += int64(len(content)) //更新偏移量
	return nil
}

// WriteFileByOffset 根据偏移量写入文件
// 传入参数：文件路径、文件内容([]byte)、偏移量
// 返回参数：错误信息
func (file *File) WriteFileByOffset(content []byte, offset int64) error {
	err := writeFileByOffset(file.file, content, offset)
	if err != nil {
		return err
	}

	//考虑并发写问题，这里不更新偏移量
	return nil
}

// AppendFile 追加文件
// 传入参数：文件路径、文件内容([]byte)
// 返回参数：错误信息
func (file *File) AppendFile(path string, content []byte) error {
	err := writeFileByOffset(file.file, content, file.offset)
	if err != nil {
		return err
	}

	file.offset += int64(len(content)) //更新偏移量
	return nil
}

// DeleteFile 删除文件
// 传入参数：文件路径
// 返回参数：错误信息
func DeleteFile(path string) error {
	return deleteFile(path)
}

// IsFileExist 判断文件是否存在
// 传入参数：文件路径
// 返回参数：bool
func IsFileExist(path string) bool {
	return isFileExist(path)
}
