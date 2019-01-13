package network

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

// TestNetwork returns back CPU tests
func TestNetwork() {
	var protocols = []string{"ip", "icmp", "icmpmsg", "tcp", "udp", "udplite"}
	x, _ := net.Connections("all")
	y, _ := net.IOCounters(true)
	z, _ := net.ProtoCounters(protocols)

	fmt.Println("====GOPSUTIL====")
	fmt.Println("Network")
	fmt.Println("Connections")
	fmt.Println(x)

	fmt.Println("counters")
	fmt.Println(y)

	fmt.Println("Protocols")
	fmt.Println(z)

}
