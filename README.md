# child-nodes-file-manager-service

约定:

- 在进行写操作之前，先调用isFileExist函数判断文件是否存在，如果不存在则调用createFile函数创建文件，存在则调用getFile函数获取文件内容。两个函数均返回一个文件对象。通过调用对象方法进行写操作。对象由中下层文件管理器保管，用后即毁。

- 在进行并发写时默认调用`WriteFileByOffset`，不允许同时调用`AppendFile`。