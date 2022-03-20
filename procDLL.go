package winapi

import "syscall"

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	sendInputProc    = user32.NewProc("SendInput")
	procSetCursorPos = user32.NewProc("SetCursorPos")
)
