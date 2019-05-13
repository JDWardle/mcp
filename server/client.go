package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/JDWardle/gocraft/protocol"
)

type Client struct {
	ID    int
	State protocol.ClientState
	conn  net.Conn
}

func NewClient(id int, conn net.Conn) *Client {
	return &Client{
		ID:    id,
		State: protocol.ClientStateHandshaking,
		conn:  conn,
	}
}

func (c *Client) HandleMessages() {
	defer c.Close()

	r := bufio.NewReader(c.conn)

	for {
		// This will block until a request is sent from the client.
		packetLength, err := protocol.ReadVarInt(r)
		if err != nil {
			fmt.Println(err)
			break
		}

		// This is more than likely a legacy server list ping so the first byte
		// was the ID of the packet.
		if c.State == protocol.ClientStateHandshaking && packetLength == 0xFE {
			ok, h := DefaultHandlers.GetHandler(c.State, packetLength)
			if !ok {
				fmt.Printf("unknown packet ID %#02x\n", packetLength)
				continue
			}

			err = h(c, r)
			if err := h(c, r); err != nil {
				fmt.Println(err)
			}
			continue
		}

		packetID, err := protocol.ReadVarInt(r)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Packet length:", packetLength)
		fmt.Printf("Packet ID: %#02x\n", packetID)

		ok, h := DefaultHandlers.GetHandler(c.State, packetID)
		if !ok {
			fmt.Printf("unknown packet ID %#02x\n", packetID)
			continue
		}

		if err := h(c, r); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (c *Client) Close() {
	fmt.Printf("client %s disconnected\n", c.conn.RemoteAddr())
	c.conn.Close()
}
