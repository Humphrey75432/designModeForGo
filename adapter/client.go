package adapter

import "fmt"

type Client struct {
}

func (c *Client) InsertLightingConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lighting connector into computer.")
	com.InsertIntoLightingPort()
}
