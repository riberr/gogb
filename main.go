package main

import "gogb/gameboy"

func main() {
	gb := gameboy.New(false)
	gb.Run()
}
