// +build !windows

package priority

import "syscall"

// priority in unix usually is a range between -19/-20 (low prio) and
// +19/+20 (high prio). 0 equals the default priority.
var priorityMapping = map[Priority]int{
	PriorityLow:    -10,
	PriorityMedium: 0,
	PriorityHigh:   10,
}

func setPriority(priority Priority) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, 0, priorityMapping[priority])
}
