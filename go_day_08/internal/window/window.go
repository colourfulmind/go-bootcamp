package window

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "application.h"
#include "window.h"
*/
import "C"
import "unsafe"

func InitApplication() {
	C.InitApplication()
}

func RunApplication() {
	C.RunApplication()
}

func CreateWindow(x, y, width, height int, title string) unsafe.Pointer {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ptr := C.Window_Create(C.int(x), C.int(y), C.int(width), C.int(height), cTitle)

	return ptr
}

func MakeKeyAndOrderFrontWindow(wndPtr unsafe.Pointer) {
	C.Window_MakeKeyAndOrderFront(wndPtr)
}
