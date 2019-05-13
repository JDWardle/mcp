package protocol

import (
	"bufio"
	"encoding/binary"
)

//go:generate stringer -type=ClientState
type ClientState int32

const (
	ClientStateHandshaking ClientState = iota
	ClientStateStatus
	ClientStateLogin
	ClientStatePlay
)

type Handshake struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       ClientState
}

func (h *Handshake) Decode(r *bufio.Reader) error {
	protocolVersion, err := ReadVarInt(r)
	if err != nil {
		return err
	}
	h.ProtocolVersion = protocolVersion

	s, err := ReadString(r)
	if err != nil {
		return err
	}
	h.ServerAddress = s

	if err := binary.Read(r, binary.BigEndian, &h.ServerPort); err != nil {
		return err
	}

	nextState, err := ReadVarInt(r)
	if err != nil {
		return err
	}
	h.NextState = ClientState(nextState)

	return nil
}
