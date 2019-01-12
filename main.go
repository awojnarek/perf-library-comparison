package main

import (
	"fmt"

	"github.com/awojnarek/perf-library-comparison/pkg/cpu"
)

func main() {
	fmt.Println("Testing CPU")
	cpu.TestCPU()
}
