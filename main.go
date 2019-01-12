package main

import (
	"github.com/awojnarek/perf-library-comparison/pkg/cpu"
	"github.com/awojnarek/perf-library-comparison/pkg/memory"
	"github.com/awojnarek/perf-library-comparison/pkg/network"
)

func main() {
	println("Testing CPU")
	cpu.TestCPU()

	println("Testing Memory")
	memory.TestMem()

	println("Testing Network")
	network.TestNetwork()
}
