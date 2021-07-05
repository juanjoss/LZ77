package lz77

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLZ77(t *testing.T) {
	suite := []string{
		"../test_files/lorem.txt",
		"../test_files/animal_farm.txt",
		"../test_files/simple_test.txt",
	}

	for _, s := range suite {
		data, err := ioutil.ReadFile(s)
		if err != nil {
			t.Errorf("error reading file %s in test.", s)
			t.FailNow()
		}

		// Initializing LZ77 struct
		lz77 := Init()

		// compressing and writing data
		compData := lz77.Compress(data)

		// decompressing data
		decompData := lz77.Decompress(compData)

		// assert
		if strings.Compare(string(data), string(decompData)) != 0 {
			t.Errorf("test %s failed: original data and decompressed data not equal.", s)
		} else {
			t.Logf("test %s passed.", s)
		}
	}
}
