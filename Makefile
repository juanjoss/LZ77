all: clean test build

test:
	cd ./LZ77 && go test -v

build:
	go build

clean:
	rm -f ./lz77
	rm -f ./test_files/simple_test.txt.lz77 ./test_files/simple_test
	rm -f ./test_files/lorem.txt.lz77 ./test_files/lorem
	rm -f ./test_files/animal_farm.txt.lz77 ./test_files/animal_farm