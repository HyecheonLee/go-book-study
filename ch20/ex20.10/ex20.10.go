package main

type Reader interface {
	Read()
}

type Closer interface {
	Closer()
}

type File struct {
}

func (f *File) Read() {

}
func ReadFile(reader Reader) {
	c := reader.(Closer)
	c.Closer()
}

func main() {
	file := &File{}
	ReadFile(file)
}
