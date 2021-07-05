package lz77

import (
	"bytes"
	"math"
)

func (lz77 *LZ77) Compress(input []byte) []byte {
	var compressedData []byte
	var nextChar byte

	// iterate over the input
	for cursor := 0; cursor < len(input); cursor++ {
		windowBeg := int(math.Max(float64(cursor-lz77.position), 0))
		bufferEnd := int(math.Min(float64(cursor+lz77.length), float64(len(input))))

		position, length := lz77.match(input[windowBeg:cursor], input[cursor:bufferEnd])

		if cursor+length >= len(input) {
			nextChar = 0
		} else {
			nextChar = input[cursor+length]
		}

		compressedData = lz77.encode(compressedData, position, length, nextChar)
		cursor += length
	}

	return compressedData
}

func (lz77 *LZ77) match(window, buffer []byte) (position, length int) {
	// iterating over buffer's lengths
	for length = len(buffer); length > 0; length-- {
		// match found in search window
		position = bytes.LastIndex(window, buffer[:length])

		if position != -1 {
			return len(window) - position, length
		}
	}

	return 0, 0
}

func (lz77 *LZ77) encode(buffer []byte, position, length int, nextChar byte) []byte {
	var buff = make([]byte, 3)

	buff[0] = byte(position & 0x00FF)
	buff[1] = byte((position&0x0F00)>>8) | byte((length&0x000F)<<4)
	buff[2] = nextChar

	return append(buffer, buff...)
}
