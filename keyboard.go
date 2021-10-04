package goautogui

/*
#include<windows.h>
*/
import "C"

const (
	KeyStatusDown uint16 = 0x8000

	KeyEventfExtendedKey = C.KEYEVENTF_EXTENDEDKEY
	KeyEventfKeyUp       = C.KEYEVENTF_KEYUP

	VkLButton          uint8 = 0x01
	VkRButton          uint8 = 0x02
	VkCancel           uint8 = 0x03
	VkMButton          uint8 = 0x04
	VkXButton1         uint8 = 0x05
	VkXButton2         uint8 = 0x06
	VkBack             uint8 = 0x08
	VkTab              uint8 = 0x09
	VkClear            uint8 = 0x0C
	VkReturn           uint8 = 0x0D
	VKShift            uint8 = 0x10
	VkControl          uint8 = 0x11
	VkMenu             uint8 = 0x12
	VkAlt              uint8 = 0x12
	VkPause            uint8 = 0x13
	VkCapital          uint8 = 0x14
	VkKana             uint8 = 0x15
	VkHangul           uint8 = 0x15
	VkIMEOn            uint8 = 0x16
	VkJunja            uint8 = 0x17
	VkFinal            uint8 = 0x18
	VkHanja            uint8 = 0x19
	VkKanji            uint8 = 0x19
	VkIMEOff           uint8 = 0x1A
	VkEscape           uint8 = 0x1B
	VkConvert          uint8 = 0x1C
	VkNonConvert       uint8 = 0x1D
	VkAccept           uint8 = 0x14
	VkModeChange       uint8 = 0x1F
	VkSpace            uint8 = 0x20
	VkPrior            uint8 = 0x21
	VkNext             uint8 = 0x22
	VkEnd              uint8 = 0x23
	VkHome             uint8 = 0x24
	VkLeft             uint8 = 0x25
	VkUp               uint8 = 0x26
	VkRight            uint8 = 0x27
	VkDown             uint8 = 0x28
	VkSelect           uint8 = 0x29
	VkPrint            uint8 = 0x2A
	VkExecute          uint8 = 0x2B
	VkSnapshot         uint8 = 0x2C
	VkInsert           uint8 = 0x2D
	VkDelete           uint8 = 0x2E
	VkHelp             uint8 = 0x2F
	Vk0                uint8 = 0x30
	Vk1                uint8 = 0x31
	Vk2                uint8 = 0x32
	Vk3                uint8 = 0x33
	Vk4                uint8 = 0x34
	Vk5                uint8 = 0x35
	Vk6                uint8 = 0x36
	Vk7                uint8 = 0x37
	Vk8                uint8 = 0x38
	Vk9                uint8 = 0x39
	VkA                uint8 = 0x41
	VkB                uint8 = 0x42
	VkC                uint8 = 0x43
	VkD                uint8 = 0x44
	VkE                uint8 = 0x45
	VkF                uint8 = 0x46
	VkG                uint8 = 0x47
	VkH                uint8 = 0x48
	VkI                uint8 = 0x49
	VkJ                uint8 = 0x4a
	VkK                uint8 = 0x4b
	VkL                uint8 = 0x4c
	VkM                uint8 = 0x4d
	VkN                uint8 = 0x4e
	VkO                uint8 = 0x4f
	VkP                uint8 = 0x50
	VkQ                uint8 = 0x51
	VkR                uint8 = 0x52
	VkS                uint8 = 0x53
	VkT                uint8 = 0x54
	VkU                uint8 = 0x55
	VkV                uint8 = 0x56
	VkW                uint8 = 0x57
	VkX                uint8 = 0x58
	VkY                uint8 = 0x59
	VkZ                uint8 = 0x5a
	VkLWin             uint8 = 0x5B
	VkRWin             uint8 = 0x5C
	VkApps             uint8 = 0x5D
	VkSleep            uint8 = 0x5F
	VkNumpad0          uint8 = 0x60
	VkNumpad1          uint8 = 0x61
	VkNumpad2          uint8 = 0x62
	VkNumpad3          uint8 = 0x63
	VkNumpad4          uint8 = 0x64
	VkNumpad5          uint8 = 0x65
	VkNumpad6          uint8 = 0x66
	VkNumpad7          uint8 = 0x67
	VkNumpad8          uint8 = 0x68
	VkNumpad9          uint8 = 0x69
	VkMultiply         uint8 = 0x6A
	VkAdd              uint8 = 0x6B
	VkSeparator        uint8 = 0x6C
	VkSubtract         uint8 = 0x6D
	VkDecimal          uint8 = 0x6E
	VkDivide           uint8 = 0x6F
	VkF1               uint8 = 0x70
	VkF2               uint8 = 0x71
	VkF3               uint8 = 0x72
	VkF4               uint8 = 0x73
	VkF5               uint8 = 0x74
	VkF6               uint8 = 0x75
	VkF7               uint8 = 0x76
	VkF8               uint8 = 0x77
	VkF9               uint8 = 0x78
	VkF10              uint8 = 0x79
	VkF11              uint8 = 0x7a
	VkF12              uint8 = 0x7b
	VkF13              uint8 = 0x7c
	VkF14              uint8 = 0x7d
	VkF15              uint8 = 0x7e
	VkF16              uint8 = 0x7f
	VkF17              uint8 = 0x80
	VkF18              uint8 = 0x81
	VkF19              uint8 = 0x82
	VkF20              uint8 = 0x83
	VkF21              uint8 = 0x84
	VkF22              uint8 = 0x85
	VkF23              uint8 = 0x86
	VkF24              uint8 = 0x87
	VkNumlock          uint8 = 0x90
	VkScroll           uint8 = 0x91
	VkLShift           uint8 = 0xA0
	VkRShift           uint8 = 0xA1
	VkLControl         uint8 = 0xA2
	VkRControl         uint8 = 0xA3
	VkLMenu            uint8 = 0xA4
	VkRMenu            uint8 = 0xA5
	VkBrowserBack      uint8 = 0xA6
	VkBrowserForward   uint8 = 0xA7
	VkBrowserRefresh   uint8 = 0xA8
	VkBrowserStop      uint8 = 0xA9
	VkBrowserSearch    uint8 = 0xAA
	VkBrowserFavorites uint8 = 0xAB
	VkBrowserHome      uint8 = 0xAC
	VkVolumeMute       uint8 = 0xAD
	VkVolumeDown       uint8 = 0xAE
	VkVolumeUp         uint8 = 0xAF
	VkMediaNextTrack   uint8 = 0xB0
)

func GetKeyState(key int) uint16 {
	ret := C.GetKeyState(C.int(key))
	return uint16(ret)
}

func KeyPressed(key int) bool {
	if GetKeyState(key)&KeyStatusDown != 0 {
		return true
	}
	return false
}

func KeyboardEvent(bVk uint8, bScan uint8, dwFlags uint32, dwExtraInfo uint64) {
	C.keybd_event(C.BYTE(bVk), C.BYTE(bScan), C.DWORD(dwFlags), C.ULONG_PTR(dwExtraInfo))
}

func PressKey(key uint8) {
	KeyboardEvent(key, 0x45, 0, 0)
}
