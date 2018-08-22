// +build windows

package priority

import (
	"fmt"
	"syscall"
)

var kernel32DLL = syscall.MustLoadDLL("kernel32")

// priority classes
var priorityMapping = map[Priority]uintptr{
	PriorityLow:    0x00004000, // equals BELOW_NORMAL_PRIORITY_CLASS
	PriorityMedium: 0x00000020, // equals NORMAL_PRIORITY_CLASS
	PriorityHigh:   0x00008000, // equals ABOVE_NORMAL_PRIORITY_CLASS
}

func setPriority(priority Priority) error {
	// https://docs.microsoft.com/en-us/windows/desktop/api/processthreadsapi/nf-processthreadsapi-setpriorityclass
	setPriorityClass, err := kernel32DLL.FindProc("SetPriorityClass")
	if err != nil {
		return err
	}
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return err
	}
	_, _, errno := syscall.Syscall(setPriorityClass.Addr(), uintptr(handle), priorityMapping[priority], 0, 0)
	if errno != 0 {
		return fmt.Errorf("SetPriorityClass(%x, %x, %x, 0x0, 0x0): %s",
			setPriorityClass.Addr(), uintptr(handle), priorityMapping[priority], errno.Error())
	}
	return nil
}
