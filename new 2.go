package main

import (
    "fmt"
    "log"
    "net"
    "strings"
)

func networkInterfaceDetails() {
	var count int

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			log.Printf("Index:%d Name:%v addr:%v, mac:%v\n", i.Index, i.Name, a, i.HardwareAddr )
		}
		count++
	}
	fmt.Println("Total interfaces : ", count)
}

func main() {
    networkInterfaceDetails()
}