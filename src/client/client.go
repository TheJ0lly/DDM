package client

import (
	"TheJ0lly/DDM/src/address"
	"TheJ0lly/DDM/src/datapack"
)

// Client is used as a sender of data.
//
// Recommended to use the CreateTyped() or CreateGeneric() methods over manual instantiation.
type Client[T any] struct {

	// Packages is the map of what data to send where.
	// The keys represent the destionation address.
	// The values represent the data to send.
	Packages map[*address.Address]*datapack.DataPack[T]
}

// CreateTyped creates a new Client with the type in brackets. This method is recommended over manual instantiation.
//
// It is equivalent to &Client[T]{Packages: make(map[address.Address]*datapack.DataPack[T])}.
func CreateTyped[T any]() *Client[T] {
	return &Client[T]{Packages: make(map[*address.Address]*datapack.DataPack[T])}
}

// CreateGeneric creates a new Client with no specific type. This method is recommended over manual instantiation.
//
// It is equivalent to &Client[any]{Packages: make(map[address.Address]*datapack.DataPack[any])}.
func CreateGeneric() *Client[any] {
	return &Client[any]{Packages: make(map[*address.Address]*datapack.DataPack[any])}
}

// AddData adds data to the corresponding DataPack, if the key exists, or creates a new DataPack otherwise.
func (c *Client[T]) AddData(dest *address.Address, data ...T) {
	if val, ok := c.Packages[dest]; ok {
		val.AddData(data...)
	} else {
		dp := datapack.CreateTyped[T]()
		c.Packages[dest] = dp
	}
}

// AddDataPack adds a whole DataPack.
func (c *Client[T]) AddDataPack(dest *address.Address, dp *datapack.DataPack[T]) {
	c.Packages[dest] = dp
}
