package filesystem

type File struct {
	Size    int64
	Content []byte
}

type Folder struct {
}

func NewFolder() *Folder {
	return &Folder{}
}

func (f *Folder) AddFile(filePath string, file File) {
}

func (f *Folder) AddFolder(folderPath string, folder *Folder) {
}

func (f *Folder) GetFolder(folderPath string) *Folder {
	return nil
}

func (f *Folder) GetFile(filePath string) *File {
	return nil
}

func (f *Folder) TotalSize() int64 {
	return 0
}
