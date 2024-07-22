package main

import (
	"fmt"

	"TheJ0lly/DDM/src/address"
	"TheJ0lly/DDM/src/datapack"
	"TheJ0lly/DDM/src/router"
)

func main() {
	r := router.Create()

	// addr, err := address.CreateNew("0.0.0.0", 8080)

	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// 	return
	// }

	// dp := datapack.Create()
	// dp.AddData(1, 2, 3, 4, 5)

	// r.AddPair(addr, dp)

	addr2, err := address.CreateNew("0.0.0.0", 6969)

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	dp2 := datapack.Create()
	dp2.AddData(7, 9, 9, 11)

	r.AddPair(addr2, dp2)

	err = r.Send()

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

}
