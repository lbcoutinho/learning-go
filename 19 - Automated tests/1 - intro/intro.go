package main

import (
	"fmt"
	"intro-tests/address"
)

func main() {
	addressType := address.AddressType("street abc")
	fmt.Println(addressType)

}
