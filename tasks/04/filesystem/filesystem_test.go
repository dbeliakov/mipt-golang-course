package filesystem

import (
	"reflect"
	"testing"
)

func TestAddFile(t *testing.T) {
	root := NewFolder()
	file := File{Size: 100, Content: []byte("content")}
	root.AddFile("subfolder1/subfolder2/file.txt", file)

	expected := &File{
		Size:    100,
		Content: []byte("content"),
	}

	if !reflect.DeepEqual(root.GetFile("subfolder1/subfolder2/file.txt"), expected) {
		t.Errorf("File was not added correctly")
	}
}

// Тест на добавление папки
func TestAddFolder(t *testing.T) {
	root := NewFolder()
	subfolder := NewFolder()
	root.AddFolder("subfolder1", subfolder)

	expected := subfolder

	actual := root.GetFolder("subfolder1")
	if expected != actual {
		t.Errorf("Folder was not added correctly")
	}
}

// Тест на получение папки по пути
func TestGetFolder(t *testing.T) {
	root := NewFolder()
	root.AddFolder("subfolder1", NewFolder())
	root.AddFolder("subfolder1/subfolder2", NewFolder())

	f := root.GetFolder("subfolder1").GetFolder("subfolder2")
	if f == nil {
		t.Errorf("Did not get the correct folder, want folder, got nil")
	}
}

// Тест на получение файла по пути
func TestGetFile(t *testing.T) {
	root := NewFolder()
	file := File{Size: 100, Content: []byte("content")}
	root.AddFile("subfolder1/subfolder2/file.txt", file)

	actual := root.GetFile("subfolder1/subfolder2/file.txt")
	if actual == nil || !reflect.DeepEqual(*actual, file) {
		t.Errorf("Did not get the correct file")
	}
}

// Тест на подсчет общего размера файлов
func TestTotalSize(t *testing.T) {
	root := NewFolder()
	root.AddFile("file1.txt", File{Size: 100})
	root.AddFile("subfolder1/file2.txt", File{Size: 200})
	root.AddFile("subfolder1/subfolder2/file3.txt", File{Size: 300})

	expected := int64(600)
	actual := root.TotalSize()
	if actual != expected {
		t.Errorf("Total size incorrect, got %d, want %d", actual, expected)
	}
}

func TestAddFolderWithNestedFolders(t *testing.T) {
	root := NewFolder()
	subSubFolder := NewFolder()
	root.AddFolder("subfolder1/subfolder2", subSubFolder)

	if f := root.GetFolder("subfolder1").GetFolder("subfolder2"); f == nil {
		t.Errorf("Nested subfolder was not added correctly")
	}

	// Проверяем также факт создания промежуточной папки subfolder1
	if f := root.GetFolder("subfolder1"); f == nil {
		t.Errorf("Intermediate folder 'subfolder1' was not created")
	}
}
