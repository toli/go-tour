package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	read, err := r.r.Read(b)
	if(err != nil) {
		return 0, err
	}
    for i:=0; i<read; i++ {
        b[i] = b[i]+13
		if b[i] > 'z' {
            b[i] = b[i] - 'z' + 'a' -1
        }
    }
	return read, nil
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
