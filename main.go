package main

import (
	"fmt"
	"io/ioutil"

	lz77 "github.com/fuato1/lz77/LZ77"
)

func main() {
	// reading file
	data, err := ioutil.ReadFile("./test_files/lorem.txt")
	if err != nil {
		panic(err)
	}

	// compressing data
	lz77 := lz77.Init()
	compressedData := lz77.Compress(data)

	// stats
	compressionRatio := float64(len(compressedData)) / float64(len(data))

	fmt.Println("Original size: ", len(data))
	fmt.Println("Compressed size: ", len(compressedData))
	fmt.Printf("Compression ratio: %.2f\n", compressionRatio)

	// writing compressed data
	err = ioutil.WriteFile("./test_files/lorem_comp", compressedData, 0644)
	if err != nil {
		panic(err)
	}

	// decompressing data
	decompressedData := lz77.Decompress(compressedData)

	// writing decompressed data
	err = ioutil.WriteFile("./test_files/lorem", decompressedData, 0644)
	if err != nil {
		panic(err)
	}
}
