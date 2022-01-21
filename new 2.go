package main
// Initiallizing packages
import (
    "fmt"
    "log"
    "net"
    "strings"
)
// Writing a function to get NIC details..
func networkInterfaceDetails() {
	var count int
// Checking for null value for interface
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}
// checking for null value for IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}

		for _, a := range addrs {
			log.Printf("Index:%d Name:%v addr:%v, mac:%v\n", i.Index, i.Name, a, i.HardwareAddr ) // Printing all the values for NIC and also NIC status
			if strings.Contains(i.Flags.String(), "up") {
				fmt.Println("Status : UP")
			} else {
				fmt.Println("Status : DOWN")
			}
		}
		count++
	}
	fmt.Println("Total interfaces : ", count)
}

func main() {
    networkInterfaceDetails()
}
