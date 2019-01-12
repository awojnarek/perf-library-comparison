package memory

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// TestMem returns back CPU tests
func TestMem() bool {

	x, _ := mem.VirtualMemory()
	y, _ := mem.SwapMemory()

	fmt.Println("====GOPSUTIL====")
	fmt.Println("Virtual Memory")
	fmt.Println(x)

	fmt.Println("Swap")
	fmt.Println(y)

	return true
}
