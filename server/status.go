package server

import (
	"bufio"
	"bytes"
	"encoding/binary"

	"github.com/JDWardle/gocraft/protocol"
)

func StatusRequestHandler(c *Client, r *bufio.Reader) error {
	// Nothing really happens with this packet.
	return nil
}

func PingHandler(c *Client, r *bufio.Reader) error {
	var pingPayload int64
	if err := binary.Read(r, binary.BigEndian, &pingPayload); err != nil {
		return err
	}

	packet := bytes.NewBuffer(protocol.VarInt(1))
	if err := binary.Write(packet, binary.BigEndian, pingPayload); err != nil {
		return err
	}
	c.conn.Write(append(protocol.VarInt(int32(len(packet.Bytes()))), packet.Bytes()...))

	return nil
}
