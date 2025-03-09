package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i := range p {
		p[i] = rot13.rot13(p[i])
	}
	return
}

func (rot13 rot13Reader) rot13(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return 'a' + (b-'a'+13)%26
	}
	if b >= 'A' && b <= 'Z' {
		return 'A' + (b-'A'+13)%26
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
