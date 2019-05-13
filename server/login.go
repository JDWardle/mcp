package server

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/JDWardle/gocraft/protocol"
)

func LoginStartHandler(c *Client, r *bufio.Reader) error {
	s, err := protocol.ReadString(r)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return errors.New("not implemented")
}

func EncryptionResponseHandler(c *Client, r *bufio.Reader) error {
	return errors.New("not implemented")
}

func LoginPluginResponseHandler(c *Client, r *bufio.Reader) error {
	return errors.New("not implemented")
}
