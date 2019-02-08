package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"runtime"
	"wlist"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var interfaceName = flag.String("i", "wlan0", "wireless interface name")
	flag.Parse()

	cells, err := wlist.Scan(*interfaceName)
	if err != nil {
		panic(err)
	}

	cellsOut, err := json.Marshal(cells)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(cellsOut))
}
