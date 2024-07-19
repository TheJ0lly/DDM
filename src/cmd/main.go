package main

import (
	"TheJ0lly/DDM/src/client"
	"TheJ0lly/DDM/src/datapack"
	"fmt"
)

func main() {
	c := client.CreateNew[int]()

	dp := datapack.CreateNew[int]()
	dp.AddData(1, 2, 3)
	c.AddDataPack("192.168.0.1:8080", dp)

	c.AddData("192.168.0.1:8080", 4, 5, 6)

	for k, v := range c.Packages {
		fmt.Printf("Key: %v --- Values: %v\n", k, v)
	}
}
