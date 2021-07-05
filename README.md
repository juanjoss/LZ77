# LZ77 Compression Algorithm.

LZ77 is a lossless compression algorithm, dictorinary based, which uses a sliding window to encode previously seen data.

This implementation is written in Go as an intention to improve my skills.

## How to test it

***Not yet.***

## How to use it
> make

    This will create the executable and clean the files generated from the tests.

> ./lz77

    Run it! 
    For now it just takes the lorem.txt in the test_files folder, and creates a compressed (lorem_comp) and a decompressed file (lorem).

## Credits

    I used this resources to understand the real implementation of LZ77. Since in reality bit operations are needed to encode and decode data.

- https://github.com/fbonhomm/LZ77
- https://www.youtube.com/watch?v=zIwTq2yPCU4&t=1297s