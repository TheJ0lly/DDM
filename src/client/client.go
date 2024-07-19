package client

import "TheJ0lly/DDM/src/datapack"

// Client is used as a sender of data
// Recommended to use the CreateNew() method or manually instantiate the Packages map.
type Client[T any] struct {

	// Packages is the map of what data to send where.
	// The keys represent the destionation address.
	// The values represent the data to send.
	Packages map[string]*datapack.DataPack[T]
}

// CreateNew has the same effect as: Client{Packages: make(map[string][]client.Package)}.
// It is a more fool-proof way to create a Client.
func CreateNew[T any]() *Client[T] {
	return &Client[T]{Packages: make(map[string]*datapack.DataPack[T])}
}

func (c *Client[T]) AddData(dest string, data ...T) {
	if val, ok := c.Packages[dest]; ok {
		val.AddData(data...)
	} else {
		dp := datapack.CreateNew[T]()
		c.Packages[dest] = dp
	}
}

func (c *Client[T]) AddDataPack(dest string, dp *datapack.DataPack[T]) {
	c.Packages[dest] = dp
}
