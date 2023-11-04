package main

import "time"

func main() {
	cartLoad()
	//fmt.Printf("%02x ", rom)

	cpu := NewCPU()

	for {
		cpu.Step()

		time.Sleep(time.Millisecond * 100)
	}
}
