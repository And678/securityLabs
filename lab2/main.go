package main

import (
	"fmt"
	"os"
	"io/ioutil"
)
func addExtension(input []byte) []byte {
	input = append(input, 0x80)
	for len(input) % 64 != 56 {
		input = append(input, 0x00)
	}
	return input
}

func addLength(input []byte, length uint64) []byte {
	for i := uint(0); i < 8; i++ {
		input = append(input, byte(length >> (8 * i)))
	}
	return input
}


func main() {
	if len(os.Args) == 1 {
		println("Usage: md5 [file] input")
	} else if len(os.Args) == 2 {
		data := []byte(os.Args[1])
		h := NewMd5Hash()
		h.message = addLength(addExtension(data), uint64(len(data)) << 3)
	} else if len(os.Args) > 2 && os.Args[1] == "file" {
		data, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		h := NewMd5Hash()
		h.message = addLength(addExtension(data), uint64(len(data)) << 3)
		fmt.Printf("%x\n\n", h.Sum())

	}
}

