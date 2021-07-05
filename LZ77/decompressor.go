package lz77

import "fmt"

func (lz77 *LZ77) Decompress(input []byte) []byte {
	var decompressedData []byte
	cursor := 0

	for cursor < len(input) {
		token := input[cursor:(cursor + 3)]

		data, nextChar := lz77.decode(decompressedData, token)

		decompressedData = append(decompressedData, data...)
		decompressedData = append(decompressedData, nextChar)

		cursor += 3
	}

	if len(decompressedData) > 0 && decompressedData[len(decompressedData)-1] == 0 {
		return decompressedData[:len(decompressedData)-1]
	}

	return decompressedData
}

func (lz77 *LZ77) decode(buffer []byte, token []byte) ([]byte, byte) {
	tmp := byteToUint32(token)

	position := len(buffer) - int(tmp&0x000FFF)
	length := position + int((tmp&0x00F000)>>12)
	nextChar := byte((tmp & 0xFF0000) >> 16)

	if position == -1 {
		fmt.Println("token: ", string(token), ", tmp: ", tmp)
	}

	return buffer[position:length], nextChar
}

func byteToUint32(data []byte) uint32 {
	var tmp uint32

	for i := 0; i < len(data); i++ {
		tmp |= uint32(data[i]) << (8 * i)
	}

	return tmp
}
