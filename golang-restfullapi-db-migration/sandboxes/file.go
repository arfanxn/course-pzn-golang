package sandboxes

import "fmt"

type File struct {
	Name string
}

func (this *File) Close() {
	fmt.Println("Close File", this.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}
