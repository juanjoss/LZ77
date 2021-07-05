package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./lorem.txt")
	if err != nil {
		panic(err)
	}

	lz77 := Init()
	compressedData := lz77.Compress(data)
	compressionRatio := float64(len(compressedData)) / float64(len(data))

	fmt.Println("Original size: ", len(data))
	fmt.Println("Compressed size: ", len(compressedData))
	fmt.Println("Compression ratio: ", compressionRatio)

	err = ioutil.WriteFile("./lorem.lz77", compressedData, 0644)
	if err != nil {
		panic(err)
	}
}
