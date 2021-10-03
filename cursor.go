package goautogui

/*
#include<windows.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

var ErrGetPosFailed = errors.New("failed to get cursor position")

type Point struct {
	X int32
	Y int32
}

func MoveTo(x, y int) (successful bool) {
	ret := C.SetCursorPos(C.int(x), C.int(y))
	if int(ret) == 1 {
		successful = true
	}
	return
}

func GetCursorPos() (*Point, error) {
	var point Point
	ret := C.GetCursorPos((*C.POINT)(unsafe.Pointer(&point)))
	if int(ret) != 1 {
		return nil, ErrGetPosFailed
	}
	return &point, nil
}
