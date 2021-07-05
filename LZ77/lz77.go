package lz77

type LZ77 struct {
	position int // 12 bits
	length   int // 4 bits
}

func Init() *LZ77 {
	return &LZ77{
		position: 4095, // 12 bits
		length:   15,   // 4 bits
	}
}
