package disk

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

// TestDisk returns back CPU tests
func TestDisk() {
	fmt.Println("====GOPSUTIL====")

	x, _ := disk.Partitions(true)
	y, _ := disk.Usage("/")
	z, _ := disk.IOCounters()

	fmt.Println("Partitions")
	fmt.Println(x)

	fmt.Println("Usage")
	fmt.Println(y)

	fmt.Println("Counters")
	fmt.Println(z)

}
