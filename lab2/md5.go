package main

import (
	"encoding/binary"
)
type basicF func(b, c, d uint32) uint32

type Md5Hash struct {
	message []byte
	buf [4]uint32
	mdx [4]uint32
}

func NewMd5Hash() Md5Hash {
	h := Md5Hash{}
	h.buf[0] = initA
	h.buf[1] = initB
	h.buf[2] = initC
	h.buf[3] = initD

	h.mdx[0] = initA
	h.mdx[1] = initB
	h.mdx[2] = initC
	h.mdx[3] = initD
	return h
}

func (h *Md5Hash) Reset() {
	h.buf[0] = h.mdx[0]
	h.buf[1] = h.mdx[1]
	h.buf[2] = h.mdx[2]
	h.buf[3] = h.mdx[3]
}


func (h *Md5Hash) runRounds() {
	// Iterate through every 64 byte (512 bit) chunk
	for c := 0; c < len(h.message); c += 64 {
		h.Reset()
		for i := 0; i < 64; i++ {
			var curFunc basicF;
			g := i
			if i < 16 {
				curFunc = funcF
			} else if i < 32 {
				curFunc = funcG
				g = (5 * i + 1) % 16
			} else if i < 48 {
				curFunc = funcH
				g = (3 * i + 5) % 16
			} else if i < 64 {
				curFunc = funcI
				g = (7 * i) % 16
			}
			h.round(curFunc, binary.LittleEndian.Uint32(h.message[c + (g * 4) : c + (g * 4) + 4]), i)
		}
		h.mdx[0] += h.buf[0]
		h.mdx[1] += h.buf[1]
		h.mdx[2] += h.buf[2]
		h.mdx[3] += h.buf[3]
	}
}

func (h *Md5Hash) Sum() [16]byte {
	h.runRounds()
	var digest [16]byte
	for i, s := range h.mdx {
		digest[i*4] = byte(s)
		digest[i*4+1] = byte(s >> 8)
		digest[i*4+2] = byte(s >> 16)
		digest[i*4+3] = byte(s >> 24)
	}
	return digest
}

func (h *Md5Hash) round(fun basicF, word uint32, i int) {
	f := h.buf[0] + fun(h.buf[1], h.buf[2], h.buf[3]) + T[i] + word
	h.buf[0] = h.buf[3]
	h.buf[3] = h.buf[2]
	h.buf[2] = h.buf[1]
	h.buf[1] = h.buf[1] + rotate(f, S[i])
}

func rotate(num uint32, rot uint32) uint32 {
	return (num << rot) | (num >> (32 - rot))
}


func funcF(b, c, d uint32) uint32 {
	return (b & c) | ((^b) & d)
}

func funcG(b, c, d uint32) uint32 {
	return (b & d) | (c & (^d))
}

func funcH(b, c, d uint32) uint32 {
	return b ^ c ^ d
}

func funcI(b, c, d uint32) uint32 {
	return c ^ (b | (^d))
}
