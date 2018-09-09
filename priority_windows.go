// +build windows

package priority

import (
	"fmt"
	"syscall"
)

// priority classes
var priorityMapping = map[Priority]uintptr{
	PriorityVeryLow:  0x00000040, // equals IDLE_PRIORITY_CLASS
	PriorityLow:      0x00004000, // equals BELOW_NORMAL_PRIORITY_CLASS
	PriorityMedium:   0x00000020, // equals NORMAL_PRIORITY_CLASS
	PriorityHigh:     0x00008000, // equals ABOVE_NORMAL_PRIORITY_CLASS
	PriorityVeryHigh: 0x00000080, // equals HIGH_PRIORITY_CLASS
}

func setPriority(priority Priority) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setPriorityClass := kernel32.NewProc("SetPriorityClass")
	if err := setPriorityClass.Find(); err != nil {
		return err
	}
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(handle)
	r1, _, errno := setPriorityClass.Call(uintptr(handle), priorityMapping[priority])
	if r1 == 0 {
		return fmt.Errorf("SetPriorityClass(0x%x, 0x%x): %s", uintptr(handle), priorityMapping[priority], errno)
	}
	return nil
}
