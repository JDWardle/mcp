package protocol

import (
	"bufio"
	"errors"
	"sync"
)

type Handler interface {
	Decode(r *bufio.Reader) error
	Encode() ([]byte, error)
}

type Packets struct {
	m map[ClientState]map[int32]Handler
	sync.RWMutex
}

func (p Packets) GetPacket(clientState ClientState, id int32) (bool, Handler) {
	defer p.RUnlock()

	if h, ok := p.m[clientState][id]; ok {
		p.RUnlock()
		return ok, h
	}

	p.RUnlock()
	return false, nil
}

var ServerPackets = Packets{
	m: map[ClientState]map[int32]Handler{
		ClientStateHandshaking: {
			0x00: HandshakeRequest{},
		},

		ClientStateStatus: {
			0x00: HandshakeRequest{},
			0x01: HandshakeRequest{},
		},
	},
}

var ClientPackets = Packets{
	m: map[ClientState]map[int32]Handler{
		ClientStateStatus: {
			0x00: HandshakeRequest{},
			0x01: HandshakeRequest{},
		},
	},
}

type HandshakeRequest struct{}

func (h HandshakeRequest) Decode(r *bufio.Reader) error {
	return errors.New("not implemented")
}

func (h HandshakeRequest) Encode() ([]byte, error) {
	return nil, errors.New("not implemented")
}

// type Packet struct {
// 	ID     int32
// 	Length int32

// 	*bufio.Reader
// }

// func NewPacket(r io.Reader) (*Packet, error) {
// 	br := bufio.NewReader(r)

// 	length, err := ReadVarInt(br)
// 	if err != nil {
// 		return nil, err
// 	}

// 	id, err := ReadVarInt(br)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Packet{
// 		ID:     id,
// 		Length: length,
// 		Reader: br,
// 	}, nil
// }

// var ClientPackets = map[ClientState]map[int32]Handler{
// 	ClientStateStatus: {
// 		// 0x00: PacketHandler{},
// 	},
// 	ClientStateLogin: {
// 		// 0x00: PacketHandler{},
// 	},
// 	ClientStatePlay: {
// 		// 0x00: PacketHandler{},
// 	},
// }

// var ServerPackets = map[ClientState]map[int32]Handler{
// 	ClientStateHandsaking: {},
// 	ClientStateStatus:     {},
// 	ClientStateLogin:      {},
// 	ClientStatePlay:       {},
// }
