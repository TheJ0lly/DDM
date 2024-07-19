package main

import (
	"TheJ0lly/DDM/src/client"
	"TheJ0lly/DDM/src/datapack"
	"fmt"
)

func main() {
	c := client.CreateNew()

	dp := datapack.CreateNew()

	for i := range 10 {
		dp.AddData(i)
	}

	c.AddDataPack("192.168.0.1:8080", dp)

	for k, v := range c.Packages {
		fmt.Printf("Key: %v --- Values: %v\n", k, v)
	}
}
