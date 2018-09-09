package priority

// Priority of a process
type Priority int

const (
	PriorityVeryLow Priority = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
	PriorityVeryHigh
)

func SetPriority(priority Priority) error {
	return setPriority(priority)
}
