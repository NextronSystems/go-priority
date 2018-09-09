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
