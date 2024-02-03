package file

// CreateDir 创建目录
// 传入参数:目录路径
// 返回参数:error
func CreateDir(path string) error {
	return createDir(path)
}

// GetFile 获取文件
// 传入参数：文件路径
// 返回参数：os.File, error
func GetFile(path string) (File, error) {
	file, err := getFile(path)
	if err != nil {
		return File{}, err
	}

	return File{File: file}, nil
}

// WriteFileByOffset 根据偏移量写入文件
// 传入参数：文件内容([]byte)、偏移量
// 返回参数：错误信息
func (file *File) WriteFileByOffset(content []byte, offset int64) error {
	err := writeFileByOffset(file.File, content, offset)
	if err != nil {
		return err
	}

	return nil
}

// AppendFile 追加文件
// 传入参数：path,文件内容([]byte)
// 返回参数：错误信息
func AppendFile(path string, content []byte) error {
	err := appendFile(path, content)
	if err != nil {
		return err
	}

	return nil
}

// AppendString 追加字符串到文本文件
// 传入参数:path,文件内容(string)
// 返回参数:错误信息
func AppendString(path, content string) error {
	err := appendString(path, content)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile 删除文件
// 传入参数：文件路径
// 返回参数：错误信息
func DeleteFile(path string) error {
	return deleteFile(path)
}

// DeleteDir 删除目录
// 传入参数:目录路径
// 返回参数:错误信息
func DeleteDir(path string) error {
	return deleteDir(path)
}

// IsPathExist 判断路径是否存在
// 传入参数：路径(目录、文件)
// 返回参数：bool
func IsPathExist(path string) bool {
	return isPathExist(path)
}

// CloseFile 关闭文件
// 传入参数:无
// 返回参数:错误信息
func (file *File) CloseFile() error {
	return file.File.Close()
}
