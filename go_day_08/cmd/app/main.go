package main

import (
	"main/internal/window"
)

// #include <stdlib.h>
import "C"

func main() {
	window.InitApplication()

	wndPtr := window.CreateWindow(400, 400, 300, 200, "School 21")
	defer C.free(wndPtr)
	window.MakeKeyAndOrderFrontWindow(wndPtr)

	window.RunApplication()
}
