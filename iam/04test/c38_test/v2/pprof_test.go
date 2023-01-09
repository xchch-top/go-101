package test

import (
	"os"
	"runtime/pprof"
	"testing"
)

func Test_Pprof(t *testing.T) {
	cpuOut, _ := os.Create("cpu.out")
	defer cpuOut.Close()
	pprof.StartCPUProfile(cpuOut)
	defer pprof.StopCPUProfile()

	memOut, _ := os.Create("mem.out")
	defer memOut.Close()
	defer pprof.WriteHeapProfile(memOut)

	Sum(3, 5)

}

func Sum(a, b int) int {
	return a + b
}
