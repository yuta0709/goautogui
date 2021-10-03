package goautogui

/*
#include<windows.h>

*/
import "C"

func GetDisplaySize() (x, y int) {
	x = int(C.GetSystemMetrics(C.SM_CXSCREEN))
	y = int(C.GetSystemMetrics(C.SM_CYSCREEN))
	return
}
