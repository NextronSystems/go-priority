package priority

import (
	"os"
	"strconv"
	"syscall"
)

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
	// On Linux, setpriority only affects the calling thread, not the whole process.
	// To work around this, we list all threads in the process and set the priority for each of them.
	// However, this is unfortunately racy as new threads can be created at any time.
	threads, err := os.ReadDir("/proc/self/task")
	if err != nil {
		return err
	}
	for _, thread := range threads {
		threadId, err := strconv.Atoi(thread.Name())
		if err != nil { // Should never happen since the task directory only contains the thread IDs
			return err
		}
		if err := syscall.Setpriority(syscall.PRIO_PROCESS, threadId, priorityMapping[priority]); err != nil {
			if err == syscall.EINVAL {
				// Bad thread ID - possibly a race where the thread terminated between the read and the setpriority
				continue
			}
			return err
		}
	}
	return nil
}
