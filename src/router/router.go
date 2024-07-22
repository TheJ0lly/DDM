package router

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"

	"TheJ0lly/DDM/src/address"
	"TheJ0lly/DDM/src/datapack"
)

// Router is a struct that holds all the destinations and the data to be delivered.
type Router struct {
	routes map[*address.Address]*datapack.DataPack
}

// Create returns an initialized instance of Router.
func Create() *Router {
	return &Router{routes: make(map[*address.Address]*datapack.DataPack)}
}

// GetDataPack returns the DataPack that is prepared to be sent to the addr.
func (r *Router) GetDataPack(addr *address.Address) (*datapack.DataPack, error) {
	if val, ok := r.routes[addr]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("did not find key: %v", addr)
}

// AddPair adds the key-value pair into the routes map.
func (r *Router) AddPair(addr *address.Address, dp *datapack.DataPack) error {
	if _, ok := r.routes[addr]; ok {
		return fmt.Errorf("key already exists")
	}
	r.routes[addr] = dp
	return nil
}

// ReplaceValue replaces the value residing at key addr.
func (r *Router) ReplaceValue(addr *address.Address, dp *datapack.DataPack) error {
	if _, ok := r.routes[addr]; ok {
		r.routes[addr] = dp
		return nil
	}

	return fmt.Errorf("key does not exist")
}

// UpdateValue updates the value currently residing at the key addr.
func (r *Router) UpdateValue(addr *address.Address, data ...int) error {
	if _, ok := r.routes[addr]; ok {
		r.routes[addr].AddData(data...)
		return nil
	}

	return fmt.Errorf("key does not exist")
}

func (r *Router) Send() error {

	ctx, cancel := context.WithCancel(context.Background())
	var result = make(chan int, len(r.routes))

	wg := sync.WaitGroup{}

	defer cancel()

	for a, dp := range r.routes {
		b, err := dp.ToJson()

		if err != nil {
			return err
		}

		conn, err := net.Dial(fmt.Sprintf("tcp%d", a.Version), a.String())

		if err != nil {
			return err
		}

		go HandleSearch(ctx, &wg, b, conn, result)
		wg.Add(1)
	}

	wg.Wait()

	return nil
}

func HandleSearch(ctx context.Context, wg *sync.WaitGroup, b []byte, conn net.Conn, result chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, err := conn.Write(b)

			if err != nil {
				fmt.Printf("ERROR: failed to write to connection: %s\n", conn.RemoteAddr())
				return
			}

			var buf = make([]byte, len(b))

			_, err = conn.Read(buf)

			if err != nil {
				fmt.Printf("ERROR: failed to read from connection: %s\n", conn.RemoteAddr())
				return
			}
			strconv.Atoi(string(buf))

			wg.Done()
			return
		}
	}
}
