package aggr_errors

import "fmt"

type AggrErrors []error

func (e AggrErrors) Error() string {
	s := ""
	for _, err := range e {
		s += fmt.Sprintf("%s\n", err.Error())
	}
	return s
}
