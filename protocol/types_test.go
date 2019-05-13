package protocol

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	tests := map[string][]byte{
		"This is a test": []byte{0x0e, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x74, 0x65, 0x73, 0x74},
		"":               []byte{0x00},
		"76435863":       []byte{0x08, 0x37, 0x36, 0x34, 0x33, 0x35, 0x38, 0x36, 0x33},
	}

	for test, expected := range tests {
		b := String(test)

		if !reflect.DeepEqual(b, expected) {
			t.Fatalf("Expected '%s' to be encoded as '%#02x' got '%#02x'", test, expected, b)
		}
		fmt.Printf("'%#02x'\n", String(test))
	}
}

func TestReadString(t *testing.T) {
	tests := map[string][]byte{
		"This is a test": []byte{0x0e, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x74, 0x65, 0x73, 0x74},
		"":               []byte{0x00},
		"76435863":       []byte{0x08, 0x37, 0x36, 0x34, 0x33, 0x35, 0x38, 0x36, 0x33},
	}

	for expected, test := range tests {
		r := bytes.NewReader(test)
		s, err := ReadString(r)
		if err != nil {
			t.Fatalf("Unexpected error: '%v'", err)
		}

		if s != expected {
			t.Fatalf("Expected '%#02x' to decode to '%s' got '%s'", test, expected, s)
		}
	}
}

func TestVarInt(t *testing.T) {
	tests := map[int32][]byte{
		0:             []byte{0x0},
		1:             []byte{0x1},
		2:             []byte{0x2},
		127:           []byte{0x7f},
		128:           []byte{0x80, 0x1},
		255:           []byte{0xff, 0x1},
		math.MaxInt32: []byte{0xff, 0xff, 0xff, 0xff, 0x07},
		-1:            []byte{0xff, 0xff, 0xff, 0xff, 0x0f},
		math.MinInt32: []byte{0x80, 0x80, 0x80, 0x80, 0x08},
	}

	for test, expected := range tests {
		b := VarInt(test)

		if len(b) == 0 {
			t.Fatalf("Expected '%d' to have one or more bytes", test)
		}

		if !reflect.DeepEqual(b, expected) {
			t.Fatalf("Expected '%d' to be encoded as '%#02x' got '%#02x'", test, expected, b)
		}
	}
}

func TestVarLong(t *testing.T) {
	tests := map[int64][]byte{
		0:             []byte{0x00},
		1:             []byte{0x01},
		2:             []byte{0x02},
		127:           []byte{0x7f},
		128:           []byte{0x80, 0x01},
		255:           []byte{0xff, 0x01},
		math.MaxInt32: []byte{0xff, 0xff, 0xff, 0xff, 0x07},
		math.MaxInt64: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
		-1:            []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01},
		math.MinInt32: []byte{0x80, 0x80, 0x80, 0x80, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01},
		math.MinInt64: []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01},
	}

	for test, expected := range tests {
		b := VarLong(test)

		if len(b) == 0 {
			t.Fatalf("Expected '%d' to have one or more bytes", test)
		}

		if !reflect.DeepEqual(b, expected) {
			t.Fatalf("Expected '%d' to be encoded as '%#02x' got '%#02x'", test, expected, b)
		}
	}
}

func TestReadVarInt(t *testing.T) {
	tests := map[int32][]byte{
		0:             []byte{0x0},
		1:             []byte{0x1},
		2:             []byte{0x2},
		127:           []byte{0x7f},
		128:           []byte{0x80, 0x1},
		255:           []byte{0xff, 0x1},
		math.MaxInt32: []byte{0xff, 0xff, 0xff, 0xff, 0x07},
		-1:            []byte{0xff, 0xff, 0xff, 0xff, 0x0f},
		math.MinInt32: []byte{0x80, 0x80, 0x80, 0x80, 0x08},
	}

	for expected, test := range tests {
		r := bytes.NewReader(test)
		v, err := ReadVarInt(r)
		if err != nil {
			t.Fatalf("Unexpected error: '%v'", err)
		}

		if v != expected {
			t.Fatalf("Expected '%#02x' to decode to '%d' got '%d'", test, expected, v)
		}
	}
}

func TestReadVarLong(t *testing.T) {
	tests := map[int64][]byte{
		0:             []byte{0x00},
		1:             []byte{0x01},
		2:             []byte{0x02},
		127:           []byte{0x7f},
		128:           []byte{0x80, 0x01},
		255:           []byte{0xff, 0x01},
		math.MaxInt32: []byte{0xff, 0xff, 0xff, 0xff, 0x07},
		math.MaxInt64: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
		-1:            []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01},
		math.MinInt32: []byte{0x80, 0x80, 0x80, 0x80, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01},
		math.MinInt64: []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01},
	}

	for expected, test := range tests {
		r := bytes.NewReader(test)
		v, err := ReadVarLong(r)
		if err != nil {
			t.Fatalf("Unexpected error: '%v'", err)
		}

		if v != expected {
			t.Fatalf("Expected '%#02x' to decode to '%d' got '%d'", test, expected, v)
		}
	}
}

func benchmarkVarInt(v int32, b *testing.B) {
	for n := 0; n < b.N; n++ {
		VarInt(v)
	}
}

func benchmarkReadVarInt(v int32, b *testing.B) {
	x := VarInt(v)
	r := bytes.NewReader(x)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ReadVarInt(r)
	}
}

func benchmarkReadVarLong(v int64, b *testing.B) {
	x := VarLong(v)
	r := bytes.NewReader(x)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ReadVarLong(r)
	}
}

func benchmarkVarLong(v int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		VarLong(v)
	}
}

func BenchmarkString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String("This is a benchmark test!")
	}
}

func BenchmarkReadString(b *testing.B) {
	bs := String("This is a benchmark test!")
	r := bytes.NewReader(bs)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ReadString(r)
	}
}

func BenchmarkVarInt0(b *testing.B)         { benchmarkVarInt(0, b) }
func BenchmarkVarInt1(b *testing.B)         { benchmarkVarInt(1, b) }
func BenchmarkVarIntNegative1(b *testing.B) { benchmarkVarInt(-1, b) }
func BenchmarkVarIntMaxInt32(b *testing.B)  { benchmarkVarInt(math.MaxInt32, b) }
func BenchmarkVarIntMinInt32(b *testing.B)  { benchmarkVarInt(math.MinInt32, b) }

func BenchmarkVarLong0(b *testing.B)         { benchmarkVarLong(0, b) }
func BenchmarkVarLong1(b *testing.B)         { benchmarkVarLong(1, b) }
func BenchmarkVarLongNegative1(b *testing.B) { benchmarkVarLong(-1, b) }
func BenchmarkVarLongMaxInt32(b *testing.B)  { benchmarkVarLong(math.MaxInt32, b) }
func BenchmarkVarLongMinInt32(b *testing.B)  { benchmarkVarLong(math.MinInt32, b) }
func BenchmarkVarLongMaxInt64(b *testing.B)  { benchmarkVarLong(math.MaxInt64, b) }
func BenchmarkVarLongMinInt64(b *testing.B)  { benchmarkVarLong(math.MinInt64, b) }

func BenchmarkReadVarInt0(b *testing.B)         { benchmarkReadVarInt(0, b) }
func BenchmarkReadVarInt1(b *testing.B)         { benchmarkReadVarInt(1, b) }
func BenchmarkReadVarIntNegative1(b *testing.B) { benchmarkReadVarInt(-1, b) }
func BenchmarkReadVarIntMaxInt32(b *testing.B)  { benchmarkReadVarInt(math.MaxInt32, b) }
func BenchmarkReadVarIntMinInt32(b *testing.B)  { benchmarkReadVarInt(math.MinInt32, b) }

func BenchmarkReadVarLong0(b *testing.B)         { benchmarkReadVarLong(0, b) }
func BenchmarkReadVarLong1(b *testing.B)         { benchmarkReadVarLong(1, b) }
func BenchmarkReadVarLongNegative1(b *testing.B) { benchmarkReadVarLong(-1, b) }
func BenchmarkReadVarLongMaxInt32(b *testing.B)  { benchmarkReadVarLong(math.MaxInt32, b) }
func BenchmarkReadVarLongMinInt32(b *testing.B)  { benchmarkReadVarLong(math.MinInt32, b) }
func BenchmarkReadVarLongMaxInt64(b *testing.B)  { benchmarkReadVarLong(math.MaxInt64, b) }
func BenchmarkReadVarLongMinInt64(b *testing.B)  { benchmarkReadVarLong(math.MinInt64, b) }
