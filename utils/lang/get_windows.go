//go:build windows

package lang

import (
	"syscall"
	"unsafe"
)

func getLanguage() (string, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLocaleName := kernel32.NewProc("GetUserDefaultLocaleName")

	var buf [85]uint16
	ret, _, err := procGetUserDefaultLocaleName.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		return "", err
	}
	return syscall.UTF16ToString(buf[:]), nil
}
