package main

import (
	"TheJ0lly/DDM/src/address"
	"TheJ0lly/DDM/src/client"
	"TheJ0lly/DDM/src/datapack"
	"fmt"
)

func main() {
	c := client.CreateGeneric()

	a, err := address.CreateNew("192.168.0.1", 8080)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	a2, err := address.CreateNew("192.168.0.2", 6969)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	dp := datapack.CreateGeneric()
	dp.AddData(1, 2, 3)
	c.AddDataPack(a, dp)
	c.AddData(a, 4, 5, 6)

	dp2 := datapack.CreateGeneric()
	dp2.AddData("Hello", "World", "From", "DDM")
	c.AddDataPack(a2, dp2)

	for k, v := range c.Packages {
		fmt.Printf("Key: %v --- Values: %v\n", k, v)
	}
}
