package child_nodes_file_manager_serivce

func InitFileManager(defaultRoot string) *FileManagerApp {
	return &FileManagerApp{
		defaultRoot: defaultRoot,
	}
}
