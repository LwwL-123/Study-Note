package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	cpunum := flag.Int("cpunum", 0, "cpunum")
	flag.Parse()
	fmt.Println("使用的cpunum:", *cpunum)
	println("系统的CPU核数",runtime.NumCPU())
	runtime.GOMAXPROCS(*cpunum)
	for i := 0; i < *cpunum - 1; i++ {
		go func() {
			for {

			}
		}()
	}
	for {

	}
}