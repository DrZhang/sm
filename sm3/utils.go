package sm3

var (
	// IV const
	IV = []uint32{
		0x7380166F, 0x4914B2B9, 0x172442D7, 0xDA8A0600, 0xA96F30BC, 0x163138AA, 0xE38DEE4D, 0xB0FB0E4E}

	// T const
	T = []uint32{
		0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519,
		0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519, 0x79CC4519,
		0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A,
		0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A,
		0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A,
		0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A,
		0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A,
		0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A, 0x7A879D8A,
	}

	// FF func map
	FF = []func(uint32, uint32, uint32) uint32{
		f1, f1, f1, f1, f1, f1, f1, f1,
		f1, f1, f1, f1, f1, f1, f1, f1,
		f2, f2, f2, f2, f2, f2, f2, f2,
		f2, f2, f2, f2, f2, f2, f2, f2,
		f2, f2, f2, f2, f2, f2, f2, f2,
		f2, f2, f2, f2, f2, f2, f2, f2,
		f2, f2, f2, f2, f2, f2, f2, f2,
		f2, f2, f2, f2, f2, f2, f2, f2,
	}

	// GG func map
	GG = []func(uint32, uint32, uint32) uint32{
		f1, f1, f1, f1, f1, f1, f1, f1,
		f1, f1, f1, f1, f1, f1, f1, f1,
		g, g, g, g, g, g, g, g,
		g, g, g, g, g, g, g, g,
		g, g, g, g, g, g, g, g,
		g, g, g, g, g, g, g, g,
		g, g, g, g, g, g, g, g,
		g, g, g, g, g, g, g, g,
	}

	// P0 func
	P0 = GenP(9, 17)

	// P1 func
	P1 = GenP(15, 23)

	// Padding bytes
	Padding = []byte{
		0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
)

// f1 func
func f1(X, Y, Z uint32) uint32 { return X ^ Y ^ Z }

// f2 func
func f2(X, Y, Z uint32) uint32 { return (X & Y) | (X & Z) | (Y & Z) }

// g func
func g(X, Y, Z uint32) uint32 { return (X & Y) | ((^X) & Z) }

// ROTL func
func ROTL(X, offset uint32) uint32 { return (X << (offset & 0x1F)) | (X >> (32 - offset&0x1F)) }

// GenP func
func GenP(offset1, offset2 uint32) func(X uint32) uint32 {
	return func(X uint32) uint32 { return X ^ ROTL(X, offset1) ^ (ROTL(X, offset2)) }
}

// ToWord func
func ToWord(i []byte) uint32 {
	return uint32(i[0])<<24 | uint32(i[1])<<16 | uint32(i[2])<<8 | uint32(i[3])<<0
}

// FromWord func
func FromWord(i uint32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}
