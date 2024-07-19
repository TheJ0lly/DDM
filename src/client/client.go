package client

import "TheJ0lly/DDM/src/datapack"

// Client is used as a sender of data.
//
// Recommended to use the CreateNew() method.
type Client[T any] struct {

	// Packages is the map of what data to send where.
	// The keys represent the destionation address.
	// The values represent the data to send.
	Packages map[string]*datapack.DataPack[T]
}

// CreateTyped creates a new Client struct with the type in brackets. This method is recommended over manual instantiation.
//
// It is equivalent to &Client[T]{Packages: make(map[string]*datapack.DataPack[T])}.
func CreateTyped[T any]() *Client[T] {
	return &Client[T]{Packages: make(map[string]*datapack.DataPack[T])}
}

func CreateGeneric() *Client[any] {
	return &Client[any]{Packages: make(map[string]*datapack.DataPack[any])}
}

func (c *Client[T]) AddData(dest string, data ...T) {
	if val, ok := c.Packages[dest]; ok {
		val.AddData(data...)
	} else {
		dp := datapack.CreateTyped[T]()
		c.Packages[dest] = dp
	}
}

func (c *Client[T]) AddDataPack(dest string, dp *datapack.DataPack[T]) {
	c.Packages[dest] = dp
}
