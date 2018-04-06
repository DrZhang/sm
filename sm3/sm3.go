package sm3

import (
	"bytes"
	"io"
)

// Sm3 struct
type Sm3 struct {
	total  [2]uint32
	state  [8]uint32
	buffer []byte
}

// NewSm3 func
func NewSm3() *Sm3 {
	return &Sm3{
		total: [2]uint32{0, 0},
		state: [8]uint32{
			IV[0], IV[1], IV[2], IV[3],
			IV[4], IV[5], IV[6], IV[7],
		},
		buffer: make([]byte, 64),
	}
}

// Update func
func (s *Sm3) Update(buf io.Reader) {
	var err error
	var count = 0
	var ilen = uint32(0)
	last := s.total[0] & 0x3F

	for {
		count, err = buf.Read(s.buffer[(last+uint32(count))&0x3F:])

		s.total[0] += uint32(count)
		if s.total[0]&0xFFFFFFFF < ilen {
			s.total[1]++
		}

		if last+uint32(count) == uint32(64) {
			// fmt.Println(hex.EncodeToString(s.buffer))
			s.process()
		}
		if err != nil && err == io.EOF {
			break
		}
	}
}

// Finish func
func (s *Sm3) Finish() []byte {
	high := s.total[0]>>29 | s.total[1]<<3
	low := s.total[0] << 3
	hb, lb := FromWord(high), FromWord(low)
	if last := s.total[0] & 0x3F; last < 56 {
		s.Update(bytes.NewReader(Padding[:56-last]))
	} else {
		s.Update(bytes.NewReader(Padding[:120-last]))
	}
	s.Update(bytes.NewReader(hb))
	s.Update(bytes.NewReader(lb))

	output := make([]byte, 32)
	for i, v := range s.state {
		for ii, b := range FromWord(v) {
			output[i*4+ii] = b
		}
	}
	return output
}

func (s *Sm3) process() {
	var ROTLA uint32
	var SS1, SS2, TT1, TT2 uint32
	var A, B, C, D uint32 = s.state[0], s.state[1], s.state[2], s.state[3]
	var E, F, G, H uint32 = s.state[4], s.state[5], s.state[6], s.state[7]
	var Wd = [64]uint32{0}
	var W = [68]uint32{0}

	for j := 0; j < 16; j++ {
		W[j] = ToWord(s.buffer[j*4 : j*4+4])
	}

	for j := 16; j < 68; j++ {
		W[j] = P1(W[j-16]^W[j-9]^ROTL(W[j-3], 15)) ^ ROTL(W[j-13], 7) ^ W[j-6]
	}

	for j := 0; j < 64; j++ {
		Wd[j] = W[j] ^ W[j+4]
	}

	for j := uint32(0); j < 64; j++ {
		ROTLA = ROTL(A, 12)
		SS1 = ROTL(ROTLA+E+ROTL(T[j], j), 7)
		SS2 = SS1 ^ ROTLA
		TT1 = FF[j](A, B, C) + D + SS2 + Wd[j]
		TT2 = GG[j](E, F, G) + H + SS1 + W[j]
		D = C
		C = ROTL(B, 9)
		B = A
		A = TT1
		H = G
		G = ROTL(F, 19)
		F = E
		E = P0(TT2)
	}
	s.state[0] ^= A
	s.state[1] ^= B
	s.state[2] ^= C
	s.state[3] ^= D
	s.state[4] ^= E
	s.state[5] ^= F
	s.state[6] ^= G
	s.state[7] ^= H
}
