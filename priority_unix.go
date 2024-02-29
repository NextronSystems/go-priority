//go:build !windows && !linux
// +build !windows,!linux

package priority

import "syscall"

// priority in unix usually is a range between -19/-20 (high prio) and
// +19/+20 (low prio). 0 equals the default priority.
var priorityMapping = map[Priority]int{
	PriorityVeryLow:  19,
	PriorityLow:      10,
	PriorityMedium:   0,
	PriorityHigh:     -10,
	PriorityVeryHigh: -20,
}

func setPriority(priority Priority) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, 0, priorityMapping[priority])
}
