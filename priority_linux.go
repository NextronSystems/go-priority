package priority

import (
	"errors"
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
	const maxThreadListings = 5

	for i := 0; i < maxThreadListings; i++ {
		var threadWithWrongPrioFound bool
		threads, err := os.ReadDir("/proc/self/task")
		if err != nil {
			return err
		}
		for _, thread := range threads {
			threadId, err := strconv.Atoi(thread.Name())
			if err != nil { // Should never happen since the task directory only contains the thread IDs
				return err
			}

			currentPriority, err := getThreadPriority(threadId)
			if err != nil {
				if err == syscall.EINVAL {
					// Bad thread ID - possibly a race where the thread terminated between the read and the getpriority
					continue
				}
				return err
			}
			if currentPriority != priorityMapping[priority] {
				threadWithWrongPrioFound = true
				if err := syscall.Setpriority(syscall.PRIO_PROCESS, threadId, priorityMapping[priority]); err != nil {
					if err == syscall.EINVAL {
						continue
					}
					return err
				}
			}
		}
		if !threadWithWrongPrioFound {
			// All threads already had the new priority when we checked them
			// Any threads they possibly spawned after the thread listing therefore also have inherited
			// the correct priority
			return nil
		}
	}

	// During each iteration before, we found new threads.
	// We give up at this point.
	return errors.New("process too volatile, could not set priority for all threads")
}

func getThreadPriority(tid int) (int, error) {
	kernelPrio, err := syscall.Getpriority(syscall.PRIO_PROCESS, tid)
	if err != nil {
		return 0, err
	}
	// getpriority returns a 0-39 range, where 39 is equivalent to priority -20 and 0 is equivalent to priority 19
	return 20 - kernelPrio, nil
}
