package file

import (
	"os"

	"github.com/238Studio/child-nodes-error-manager/errpack"
)

// 创建目录
// 传入参数:目录路径
// 返回参数:error
func createDir(path string) error {
	err := os.Mkdir(path, 0744)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}

	return nil
}

// 获取文件
// 传入参数：文件路径
// 返回参数：os.File, error
func getFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		return nil, errpack.NewError(errpack.File, errpack.TrivialException, err)
	}
	return file, nil
}

// 打开文件并在尾部追加数据
// 传入参数:文件路径，文件内容([]byte)
// 返回参数:error
func appendFile(path string, content []byte) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}

	_, err = file.Write(content)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}

	return nil
}

// 根据偏移量写入文件
// 传入:*os.file,文件内容([]byte),偏移量
// 传出:错误信息
func writeFileByOffset(file *os.File, content []byte, offset int64) error {
	_, err := file.WriteAt(content, offset)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}
	return nil
}

// 根据偏移量写字符串
// 传入:path,content(string)
// 传出:错误信息
func appendString(path, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}

	return nil
}

// 删除文件
// 传入参数：文件路径
// 返回参数：错误信息
func deleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}
	return nil
}

// 删除目录
// 传入参数:目录路径
// 返回参数:错误信息
func deleteDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return errpack.NewError(errpack.File, errpack.TrivialException, err)
	}
	return nil
}

// 判断路径是否存在
// 传入参数：路径(目录、文件)
// 返回参数：bool
func isPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
