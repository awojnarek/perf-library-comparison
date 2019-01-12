package network

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

// TestNetwork returns back CPU tests
func TestNetwork() {

	x, _ := net.Connections("all")

	fmt.Println("====GOPSUTIL====")
	fmt.Println("Network")
	fmt.Println(x)

}
