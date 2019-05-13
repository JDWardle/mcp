package server

import (
	"bufio"
	"errors"

	"github.com/JDWardle/gocraft/protocol"
)

func HandshakeHandler(c *Client, r *bufio.Reader) error {
	h := &protocol.Handshake{}
	err := h.Decode(r)
	if err != nil {
		return err
	}

	c.State = h.NextState

	if h.NextState == protocol.ClientStateStatus {
		c.State = h.NextState
		packet := append(protocol.VarInt(0), protocol.String(`{"version":{"name":"1.13.2","protocol":404},"players":{"max":100000000,"online":0,"sample":[]},"description":{"text":"Hello Minecraft from Go!"}}`)...)
		packet = append(protocol.VarInt(int32(len(packet))), packet...)
		c.conn.Write(packet)
	}

	return nil
}

func LegacyServerListPingHandler(c *Client, r *bufio.Reader) error {
	return errors.New("not implemented")
}
