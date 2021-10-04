package goautogui

/*
#include<windows.h>
*/
import "C"
import (
	"unsafe"
)

type MBType uint32
type MBResult int32

const (
	MBOK                  MBType = C.MB_OK
	MBOKCancel            MBType = C.MB_OKCANCEL
	MBAbortRetryIgnore    MBType = C.MB_ABORTRETRYIGNORE
	MBYesNoCancel         MBType = C.MB_YESNOCANCEL
	MBYesNo               MBType = C.MB_YESNO
	MBRetryCancel         MBType = C.MB_RETRYCANCEL
	MBCancelRetryContinue MBType = C.MB_CANCELTRYCONTINUE

	MBIconHand        MBType = C.MB_ICONHAND
	MBIconStop        MBType = C.MB_ICONSTOP
	MBIconError       MBType = C.MB_ICONERROR
	MBIconQuestion    MBType = C.MB_ICONQUESTION
	MBIconExclamation MBType = C.MB_ICONEXCLAMATION
	MBIconWarning     MBType = C.MB_ICONWARNING
	MBIconAsterisk    MBType = C.MB_ICONASTERISK
	MBIconInformation MBType = C.MB_ICONINFORMATION

	IDOK       MBResult = C.IDOK
	IDCancel   MBResult = C.IDCANCEL
	IDAbort    MBResult = C.IDABORT
	IDRetry    MBResult = C.IDRETRY
	IDIgnore   MBResult = C.IDIGNORE
	IDYes      MBResult = C.IDYES
	IDNo       MBResult = C.IDNO
	IDTryAgain MBResult = C.IDTRYAGAIN
	IDContinue MBResult = C.IDCONTINUE
)

func MessageBox(text string, caption string, button MBType, icon MBType) MBResult {
	t := C.CString(text)
	c := C.CString(caption)
	ret := C.MessageBox(nil, t, c, C.UINT(button|icon))
	C.free(unsafe.Pointer(t))
	C.free(unsafe.Pointer(c))
	return MBResult(ret)
}

func Alert(text string, caption string){
	MessageBox(text, caption, MBOK, MBIconInformation)
}

func Confirm(text string, caption string)bool{
	result := MessageBox(text, caption, MBOKCancel, MBIconInformation)
	switch result{
	case IDOK:
		return true
	default:
		return false
	}
}
