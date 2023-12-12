package child_nodes_file_manager_serivce

// GetRootPath 获取一个根路径
// 传入：名
// 传出：路径
func (manager *FileManagerApp) GetRootPath(name string) string {
	return manager.defaultRoot + "/" + name
}

// GetPath 获取一个路径
// 传入：相对路径
// 传出：路径
func (manager *FileManagerApp) GetPath(relatedPath []string) string {
	path := manager.defaultRoot
	for _, p := range relatedPath {
		path += "/" + p
	}
	return path
}
