package sandboxes

import "fmt"

type Connection struct {
	*File
}

func (this *Connection) Close() {
	fmt.Println("Close connection", this.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	conn := &Connection{File: file}
	return conn, func() {
		conn.Close()
	}
}
