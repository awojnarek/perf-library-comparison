package load

import (
	"fmt"

	"github.com/shirou/gopsutil/load"
)

// TestLoad returns back CPU tests
func TestLoad() {
	fmt.Println("====GOPSUTIL====")

	x, _ := load.Avg()

	fmt.Println("Load")
	fmt.Println(x)
}
