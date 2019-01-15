package process

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

// TestProcess returns back CPU tests
func TestProcess() {

	fmt.Println("====GOPSUTIL====")
	fmt.Println("Processes")

	x, _ := process.Processes()

	for _, y := range x {

		/*
			ioc, _ := y.IOCounters()
			fmt.Println(y, "IOCounters", ioc)
		*/

		ioc, _ := y.MemoryMaps(true)
		fmt.Println(y, "Memory Maps", ioc)

		/*
			noc, _ := y.NetIOCounters(true)
			thr, _ := y.NumThreads()
			fdr, _ := y.NumFDs()
			mem, _ := y.MemoryInfo()
			memp, _ := y.MemoryPercent()
		*/

	}
}
