# LZ77 Compression Algorithm.

LZ77 is a text compression algorithm, dictorinary based, which uses a sliding window to encode previously seen data.

This implementation is written in Go as an intention to improve my skills.

## How to test it

***Not yet.***

## How to use it

    This will create the executable and clean the files generated from the test.
> make

    Run it! 
    For now it just takes the lorem.txt in the test_files folder, and creates a compressed (lorem_comp) and a decompressed file (lorem).
> ./lz77