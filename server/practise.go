package main

import (
	"bytes"
	"fmt"
)

func bufferRead() {
	var bf bytes.Buffer
	bf.Write([]byte("hello"))
	fmt.Println("bf", bf)

	b1Str := bytes.NewBufferString("b1Str")
	b2 := bytes.NewBuffer([]byte{'b', '2'})
	b3 := bytes.NewBuffer([]byte("b3"))
	fmt.Println("b1", b1Str)
	fmt.Println("b2", b2)
	fmt.Println("b3", b3)

	by := []byte("byte")
	b2.Write(by)
	fmt.Println("by", b2)
	bf.Write(by)
	fmt.Println("bf", bf)

	barr := bf.Next(1)
	binf := bytes.NewBuffer(barr)
	fmt.Println("binf", binf.String())
}

type UnError struct {
	Msg string
}

func (u UnError) Error() string {
	return "afweb"
}

func tesErr() error {
	return UnError{"as"}
}
func main() {
	bufferRead()
}
