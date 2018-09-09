package main

import (
	"crypto/md5"
	"fmt"
)
func addExtension(input []byte) []byte {
	input = append(input, 0x80)
	for len(input) % 64 != 56 {
		input = append(input, 0x00)
	}
	return input
}

func addLength(input []byte, length uint64) []byte {
	// Length in bits (may need bigger type)
	for i := uint(0); i < 8; i++ {
		input = append(input, byte(length >> (8 * i)))
	}
	return input
}


func main() {
	h := newMd5Hash()
	data := []byte("abc2")
	h.message = addLength(addExtension(data), uint64(len(data)) << 3)


	fmt.Printf("\n%x\n\n", md5.Sum(data))
	
	fmt.Printf("%x\n\n", h.runRounds())
	/*fmt.Printf("%X\n",h.mdx[0])
	fmt.Printf("%X\n",h.mdx[1])
	fmt.Printf("%X\n",h.mdx[2])
	fmt.Printf("%X\n\n",h.mdx[3])*/
}

