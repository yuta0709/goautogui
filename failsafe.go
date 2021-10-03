package goautogui

/*
#include<windows.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

var ErrFailSafe = errors.New("fail safe triggerd")

func FailSafeCheck(x, y int32) {
	var point Point
	successful := C.GetCursorPos((*C.POINT)(unsafe.Pointer(&point)))
	if successful == 1 {
		if point.X == 0 && point.Y == 0 {
			panic(ErrFailSafe)
		}
	}
}
func DefaultFailSafeCheck() {
	FailSafeCheck(0, 0)
}
