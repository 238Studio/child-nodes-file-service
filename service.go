package file

import (
	"os"

	_const "github.com/238Studio/child-nodes-assist/const"
	"github.com/238Studio/child-nodes-assist/util"
)

// 获取文件
// 传入参数：文件路径
// 返回参数：os.File, error
func getFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, util.NewError(_const.File, _const.TrivialException, err)
	}
	return file, nil
}

// 创建文件
// 传入:文件路径
// 传出:错误信息
func createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, util.NewError(_const.File, _const.TrivialException, err)
	}
	return file, nil
}

// 根据偏移量写入文件
// 传入:*os.file,文件内容([]byte),偏移量
// 传出:错误信息
func writeFileByOffset(file *os.File, content []byte, offset int64) error {
	_, err := file.WriteAt(content, offset)
	if err != nil {
		return util.NewError(_const.File, _const.TrivialException, err)
	}
	return nil
}

// 删除文件
// 传入参数：文件路径
// 返回参数：错误信息
func deleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return util.NewError(_const.File, _const.TrivialException, err)
	}
	return nil
}

// 判断文件是否存在
// 传入参数：文件路径
// 返回参数：bool
func isFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
