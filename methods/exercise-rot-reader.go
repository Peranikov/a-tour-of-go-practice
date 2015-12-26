package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}


//ABCDEFGHIJKLMNOPQRSTUVWXYZ
//NOPQRSTUVWXYZABCDEFGHIJKLM

func rot13(sb byte) byte {
	s := rune(sb)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		return sb + 13
	}
	if s >= 'n' && s <= 'z' || s >= 'n' && s <= 'Z' {
		return sb - 13
	}

	return sb
}

func (reader rot13Reader) Read(b []byte) (num int, err error) {
	num, err = reader.r.Read(b)
	for i := 0; i < num; i++ {
		b[i] = rot13(b[i])
	}
	return
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
