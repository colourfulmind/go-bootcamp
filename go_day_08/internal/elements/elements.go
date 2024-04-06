package elements

import (
	"errors"
	"fmt"
	"unsafe"
)

func GetElement(arr []int, idx int) (int, error) {
	err := CheckInput(arr, idx)
	if err != nil {
		return 0, err
	}

	p := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(idx)*unsafe.Sizeof(arr[0]))

	return *(*int)(p), nil
}

func CheckInput(arr []int, idx int) error {
	if arr == nil || len(arr) == 0 {
		return errors.New(fmt.Sprintf("invalid data provided: %v", arr))
	}
	if idx > len(arr) || idx < 0 {
		return errors.New(fmt.Sprintf("index out of range: %d\n", idx))
	}
	return nil
}
