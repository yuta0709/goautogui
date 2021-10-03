package goautogui

/*
#include<windows.h>
*/
import "C"

type MouseEventFlag C.DWORD

const (
	EventAbsolute   MouseEventFlag = C.MOUSEEVENTF_ABSOLUTE
	EventLeftDown   MouseEventFlag = C.MOUSEEVENTF_LEFTDOWN
	EventLeftUp     MouseEventFlag = C.MOUSEEVENTF_LEFTUP
	EventMiddleDown MouseEventFlag = C.MOUSEEVENTF_MIDDLEDOWN
	EventMiddleUp   MouseEventFlag = C.MOUSEEVENTF_MIDDLEUP
	EventMove       MouseEventFlag = C.MOUSEEVENTF_MOVE
	EventRightDown  MouseEventFlag = C.MOUSEEVENTF_RIGHTDOWN
	EventRightUp    MouseEventFlag = C.MOUSEEVENTF_RIGHTUP
	EventWheel      MouseEventFlag = C.MOUSEEVENTF_WHEEL
	EventXDown      MouseEventFlag = C.MOUSEEVENTF_XDOWN
	EventXUp        MouseEventFlag = C.MOUSEEVENTF_XUP
	//EventHWheel     MouseEventFlag = C.MOUSEEVENTF_HWHEEL
)

func MouseEvent(e MouseEventFlag, dx uint32, dy uint32, dw uint32, dwExtra uint64) {
	C.mouse_event(C.DWORD(e), C.DWORD(dx), C.DWORD(dy), C.DWORD(dw), C.ULONG_PTR(dwExtra))
}

func Click() {
	MouseEvent(EventLeftDown, 0, 0, 0, 0)
	MouseEvent(EventLeftUp, 0, 0, 0, 0)
}
func DoubleClick() {
	Click()
	Click()
}
