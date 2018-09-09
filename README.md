# go-priority

Manipule the priority of the GO program

## Example

```golang
import (
    priority "github.com/Codehardt/go-priority"
)

func main() {
    if err := priority.SetPriority(priority.PriorityLow); err != nil {
        ...
    }
    ...
}
```

## Priorities

| Priority | [Unix](https://linux.die.net/man/1/nice) | [Windows](https://docs.microsoft.com/en-us/windows/desktop/procthread/scheduling-priorities) |
| - | - | - |
| PriorityVeryLow | 19 | IDLE_PRIORITY_CLASS |
| PriorityLow | 10 | BELOW_NORMAL_PRIORITY_CLASS | 
| PriorityMedium | 0 | NORMAL_PRIORITY_CLASS |
| PriorityHigh | -10 | ABOVE_NORMAL_PRIORITY_CLASS | 
| PriorityVeryHigh | -20 | HIGH_PRIORITY_CLASS |
