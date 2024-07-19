package client

import "TheJ0lly/DDM/src/datapack"

// Client is used as a sender of data
// Recommended to use the CreateNew() method or manually instantiate the Packages map.
type Client struct {

	// Packages is the map of what data to send where.
	// The keys represent the destionation address.
	// The values represent the data to send.
	Packages map[string]*datapack.DataPack
}

// CreateNew has the same effect as: Client{Packages: make(map[string][]client.Package)}.
// It is a more fool-proof way to create a Client.
func CreateNew() *Client {
	return &Client{Packages: make(map[string]*datapack.DataPack)}
}

func (c *Client) AddDataPack(dest string, dp *datapack.DataPack) {
	c.Packages[dest] = dp
}
