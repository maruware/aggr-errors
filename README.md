# aggr-errors

## Usage

```go
import "github.com/maruware/aggr-errors"

func sample() error {
    errors := aggr_errors.AggrErrors{}
    if err := process1(); err != nil {
        errors = append(errors, err)
    }

    if err := process2(); err != nil {
        errors = append(errors, err)
    }
    return errors
}
```