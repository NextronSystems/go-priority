package priority

// Priority of a process
type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

func SetPriority(priority Priority) error {
	return setPriority(priority)
}
