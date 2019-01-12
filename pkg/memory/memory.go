package memory

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// TestMem returns back CPU tests
func TestMem() bool {

	x, _ := mem.VirtualMemory()

	fmt.Println("====GOPSUTIL====")
	fmt.Println(x)

	return true
}
