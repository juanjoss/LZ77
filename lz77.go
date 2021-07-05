package main

import (
	"bytes"
	"math"
)

type LZ77 struct {
	position int // 12 bits
	length   int // 4 bits
}

func Init() *LZ77 {
	return &LZ77{
		position: 4095,
		length:   15,
	}
}

func (lz77 *LZ77) Compress(input []byte) []byte {
	var compressedData []byte
	var nextChar byte
	cursor := 0

	// iterate over the input
	for cursor < len(input) {
		windowBeg := int(math.Max(float64(cursor-lz77.position), 0))
		bufferEnd := int(math.Min(float64(cursor+lz77.length), float64(len(input))))

		position, length := lz77.match(input[windowBeg:cursor], input[cursor:bufferEnd])

		if cursor+length >= len(input) {
			nextChar = 0
		} else {
			nextChar = input[cursor+length]
		}

		compressedData = lz77.encode(compressedData, position, length, nextChar)
		cursor += length + 1
	}

	return compressedData
}

func (lz77 *LZ77) match(window, buffer []byte) (position, length int) {
	// iterating over buffer's lengths
	for length := len(buffer); length > 0; length-- {
		// match found in search window
		position = bytes.LastIndex(window, buffer[:length])

		if position != -1 {
			return len(buffer) - position, length
		}
	}

	return 0, 0
}

func (lz77 *LZ77) encode(compressData []byte, position, length int, nextChar byte) []byte {
	buff := make([]byte, 3)

	buff[0] = byte(position & 0x00FF) // 0x00FF = 255
	// 0x0F00 = 3840, 0x000F = 15
	buff[1] = byte((position&0x0F00)>>8) | byte((length&0x000F)<<4)
	buff[2] = nextChar

	return append(compressData, buff...)
}
