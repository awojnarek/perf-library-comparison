package cpu

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

// TestCPU returns back CPU tests
func TestCPU() bool {

	x, _ := cpu.Info()
	y, _ := cpu.Times(false)

	fmt.Println("====GOPSUTIL====")
	fmt.Println("Configs")
	fmt.Println(x)

	fmt.Println("Metrics")
	fmt.Println(y)
	return true
}
