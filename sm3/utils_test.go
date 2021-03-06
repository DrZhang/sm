package sm3

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
)

func TestP(t *testing.T) {
	var i uint32 = 0xFFFFFFF0
	P1523 := GenP(15, 23)
	t.Fatalf("\n%0.8x\n%0.8x\n", i, P1523(i))
}

func TestFromWord(t *testing.T) {
	var i uint32 = 0x01020408
	t.Fatal(FromWord(i))
}

func TestToWord(t *testing.T) {
	var i uint32 = 0x01020408
	t.Fatalf("%v", ToWord(FromWord(i)))
}

func TestLoadBytes(t *testing.T) {
	data := bytes.NewReader([]byte{0x01, 0x02, 0x04, 0x08, 0x0C, 0x0F})
	var buf [4]byte
	for {
		buf[3] = 0xFF
		_, err := data.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				return
			}
			t.Fatalf(err.Error(), hex.EncodeToString(buf[:]))
		}
		t.Error(hex.EncodeToString(buf[:]))
	}
}

func TestBytes(t *testing.T) {
	data := []byte{0x01, 0x02, 0x04, 0x08, 0x0C, 0x0F}
	t.Error(data[:6])
}

func TestSm3abc(t *testing.T) {
	s := NewSm3()
	data := []byte{0x61, 0x62, 0x63}
	s.Update(bytes.NewReader(data))
	t.Error(hex.EncodeToString(s.Finish()))
}

func TestSm3abcd512(t *testing.T) {
	s := NewSm3()
	data := []byte{
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
		0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64,
	}
	s.Update(bytes.NewReader(data))
	t.Error(hex.EncodeToString(s.Finish()))
}
