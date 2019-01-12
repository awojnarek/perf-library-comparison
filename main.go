package main

import (
	"github.com/awojnarek/perf-library-comparison/pkg/cpu"
	"github.com/awojnarek/perf-library-comparison/pkg/memory"
)

func main() {
	println("Testing CPU")
	cpu.TestCPU()

	println("Testing Memory")
	memory.TestMem()
}
