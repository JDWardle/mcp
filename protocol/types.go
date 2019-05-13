package protocol

import (
	"encoding/binary"
	"errors"
	"io"
)

// String encodes the passed in string as it's byte representation prefixed with
// a VarInt of the total length of the string.
// See https://wiki.vg/Protocol for more info.
func String(s string) []byte {
	byteString := []byte(s)
	b := VarInt(int32(len(byteString)))
	return append(b, byteString...)
}

// ReadString reads bytes from a ByteReader and returns the string
// representation of the bytes that were read. Strings are prefixed with the
// length of the expected string as a VarInt.
// See https://wiki.vg/Protocol for more info.
func ReadString(r io.ByteReader) (string, error) {
	n, err := ReadVarInt(r)
	if err != nil {
		return "", err
	}

	if n <= 0 {
		return "", nil
	}

	s := make([]byte, n)
	for i := range s {
		b, err := r.ReadByte()
		if err != nil {
			return "", err
		}
		s[i] = b
	}

	return string(s), nil
}

// SizeVarInt returns the size in bytes that is needed to represent an int32 as
// a VarInt.
func SizeVarInt(x uint32) int {
	for i := uint(1); i < binary.MaxVarintLen32; i++ {
		if x < 1<<(7*i) {
			return int(i)
		}
	}

	return binary.MaxVarintLen32
}

// SizeVarLong returns the size in bytes that is neeeded to represent an int64
// as a VarLong.
func SizeVarLong(x uint64) int {
	for i := uint(1); i < binary.MaxVarintLen64; i++ {
		if x < 1<<(7*i) {
			return int(i)
		}
	}

	return binary.MaxVarintLen64
}

// VarInt encodes int32 numbers into it's binary variable representation. Uses
// at most 5 bytes to represent any value between MinInt32 and MaxInt32.
// Negative numbers always use the maximum number of bytes.
// See https://wiki.vg/Protocol#VarInt_and_VarLong for more info.
func VarInt(v int32) []byte {
	var n int

	// This will convert negative int32 values to MaxUint32 + -v
	// -1 becomes MaxUint32
	// MinInt32 becomes MaxInt32 + 1
	x := uint32(v)

	buf := make([]byte, SizeVarInt(x))
	for {
		b := byte(x & 0x7F)
		x >>= 7

		if x != 0 {
			buf[n] = b | 0x80
			n++
		} else {
			buf[n] = b
			n++
			break
		}
	}

	return buf
}

// VarLong encodes int64 numbers into it's binary variable representation. Uses
// at most 10 bytes to represent any value between MinInt64 and MaxInt64.
// Negative numbers always use the maximum number of bytes.
// See https://wiki.vg/Protocol#VarInt_and_VarLong for more info.
func VarLong(v int64) []byte {
	var n int

	// This will convert negative int64 values to MaxUint64 + -v
	// -1 becomes MaxUint64
	// MinInt64 becomes MaxInt64 + 1
	x := uint64(v)

	buf := make([]byte, SizeVarLong(x))
	for {
		b := byte(x & 0x7F)
		x >>= 7

		if x != 0 {
			buf[n] = b | 0x80
			n++
		} else {
			buf[n] = b
			n++
			break
		}
	}

	return buf
}

// ReadVarInt reads bytes from a ByteReader until it reaches the end of a VarInt.
// Returns an error if the VarInt ends up being larger than 5 bytes or EOF is
// returned from the reader.
func ReadVarInt(r io.ByteReader) (int32, error) {
	var res uint32
	var n uint

	for {
		r, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		v := uint32(r & 0x7f)
		res |= (v << (7 * n))

		n++
		if n > 5 {
			return 0, errors.New("invalid VarInt format")
		}

		if (r & 0x80) == 0 {
			break
		}
	}

	return int32(res), nil
}

// ReadVarLong reads bytes from a ByteReader until it reaches the end of a
// VarLong. Returns an error if the VarLong ends up being larger than 10 bytes
// or EOF is returned from the reader.
func ReadVarLong(r io.ByteReader) (int64, error) {
	var res uint64
	var n uint

	for {
		r, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		v := uint64(r & 0x7f)
		res |= (v << (7 * n))

		n++
		if n > 10 {
			return 0, errors.New("invalid VarLong format")
		}

		if (r & 0x80) == 0 {
			break
		}
	}

	return int64(res), nil
}

// Following encoding/xml for implementing the Protocol marshaler/unmarshaler

// type VarInt int32

// func (v VarInt) MarshalProtocol(e *Encoder) error {
// 	return errors.New("not implemented")
// }

// func (v *VarInt) UnmarshalProtocol(d *Decoder) error {
// 	return errors.New("not implemented")
// }

// type String string

// func (s String) MarshalProtocol(e *Encoder) error {
// 	return errors.New("not implemented")
// }

// func (s *String) UnmarshalProtocol(d *Decoder) error {
// 	return errors.New("not implemented")
// }

// type Decoder struct {
// 	io.ByteReader
// }

// func NewDecoder(r io.Reader) *Decoder {
// 	d := &Decoder{}
// 	if rb, ok := r.(io.ByteReader); ok {
// 		d.ByteReader = rb
// 	} else {
// 		d.ByteReader = bufio.NewReader(r)
// 	}
// 	return d
// }

// type Encoder struct {
// 	io.ByteWriter
// }

// func NewWriter(w io.Writer) *Encoder {
// 	return &Encoder{
// 		bufio.NewWriter(w),
// 	}
// }

// type Marshaler interface {
// 	MarshalProtocol(e *Encoder) error
// }

// type Unmarshaler interface {
// 	UnmarshalProtocol(d *Decoder) error
// }

// type ProtocolID int32

// type String string

// func ReadString(r io.ByteReader) (string, error) {
// 	n, err := binary.ReadUvarint(r)
// 	if err != nil {
// 		return "", nil
// 	}

// 	if n == 0 {
// 		return "", nil
// 	}

// 	s := make([]byte, n)
// 	for i := range s {
// 		b, err := r.ReadByte()
// 		if err != nil {
// 			return "", err
// 		}
// 		s[i] = b
// 	}

// 	return string(s), nil
// }

// func PutString(r *bytes.Reader) error {

// }

// Should be handled:

// Boolean bool binary.Read
// Byte int8 binary.Read
// UnsignedByte uint8 binary.Read
// Short int16 binary.Read
// UnsignedShort uint16 binary.Read
// Int int32 binary.Read
// Long int64 binary.Read
// Angle int8 binary.Read

// Handled with custom implementation:
// String(n) string protocol.String protocol.ReadString
// VarInt int32 protocol.VarInt protocol.ReadVarInt
// VarLong int64 protocol.VarLong protocol.ReadVarLong

// Needs custom implementation:
// Float float32 custom? or binary.Read
// Double float64 custom? or binary.Read
// Chat JSON custom lenient decoding? max length 32767 (bytes? bits? characters?)
// Identifier string custom max length 32767 (bytes? bits? characters?)
// EntityMetadata []byte custom https://wiki.vg/Entity_metadata#Entity_Metadata_Format
// Slot []byte https://wiki.vg/Slot_Data
// NBTTag []byte https://wiki.vg/NBT
// Position int64 custom https://wiki.vg/Protocol#Position
// UUID string custom 128 bit integer
// Optional X custom
// Array of X custom
// X Enum custom
// Byte Array []byte custom
