all:
	clean build

build:
	go build

clean:
	rm -f ./lz77
	rm -f ./test_files/data_comp ./test_files/data
	rm -f ./test_files/lorem_comp ./test_files/lorem