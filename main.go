package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	lz77 "github.com/fuato1/lz77/LZ77"
)

func main() {
	c := flag.String("c", "", "compress file.")
	d := flag.String("d", "", "decompress file.")

	flag.Parse()

	lz77 := lz77.Init()

	if *c != "" {
		// reading file
		fmt.Println("filepath: ", *c)
		data, err := ioutil.ReadFile(*c)
		if err != nil {
			panic(err)
		}

		// compressing data
		compressedData := lz77.Compress(data)

		// stats
		compressionRatio := float64(len(compressedData)) / float64(len(data))

		fmt.Println("Original size: ", len(data))
		fmt.Println("Compressed size: ", len(compressedData))
		fmt.Printf("Compression ratio: %.2f\n", compressionRatio)

		// writing compressed data
		err = ioutil.WriteFile(*c+".lz77", compressedData, 0644)
		if err != nil {
			panic(err)
		}
	} else if *d != "" {
		// reading file
		fmt.Println("filepath: ", *d)
		compData, err := ioutil.ReadFile(*d)
		if err != nil {
			panic("error in decompressor: couldn't open filepath.")
		}

		if filepath.Ext(*d) != ".lz77" {
			panic("decompressor requires a .lz77 file.")
		}
		file := strings.Split(*d, filepath.Ext(*d))[0]
		file = strings.Split(file, filepath.Ext(file))[0]

		// decompressing data
		decompressedData := lz77.Decompress(compData)

		// writing decompressed data
		fmt.Println("output: ", file)
		err = ioutil.WriteFile(file, decompressedData, 0644)
		if err != nil {
			panic("error in decompressor: couldn't write file.")
		}
	} else {
		fmt.Print("\nusage: ./lz77 [-c=filepath | -d=filepath]\n")
	}
}
