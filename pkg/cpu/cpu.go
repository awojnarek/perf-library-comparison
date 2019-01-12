package cpu

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

// TestCPU returns back CPU tests
func TestCPU() bool {

	x, _ := cpu.Info()

	fmt.Println(x)
	return true
}
