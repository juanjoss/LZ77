# LZ77 Compression Algorithm.

LZ77 is a lossless compression algorithm, dictorinary based, which uses a sliding window to encode previously seen data.

This implementation is written in Go as an intention to improve my skills.

## How to build and test it

The makefile cleans the files created if you tested with files from /test_files, runs the test and creates the executable.

> make

## How to use it

Run it!

The executable takes a flag (-c or -d) along with the path to the file to compress/decompress.

> ./lz77 [-c=filepath | -d=filepath]

## Credits

I used this resources to understand how bit operations are used in LZ77 to encode and decode the data.

- https://github.com/fbonhomm/LZ77
- https://www.youtube.com/watch?v=zIwTq2yPCU4&t=1297s