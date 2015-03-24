package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FieldErrors map[string][]error

// Error implements error
func (f FieldErrors) Error() string {
	s := []string{
		"errors:",
	}

	for k, e := range f {
		var fs []string
		for _, v := range e {
			fs = append(fs, fmt.Sprintf(fmt.Sprintf("    * %s", v)))
		}
		s = append(s, fmt.Sprintf("  - %s\n%s", k, strings.Join(fs, "\n")))
	}

	return strings.Join(s, "\n")
}

func (f FieldErrors) MarshalJSON() ([]byte, error) {
	m := make(map[string][]string)
	for k, e := range f {
		for _, v := range e {
			m[k] = append(m[k], v.Error())
		}
	}

	if len(m) == 0 {
		return nil, nil
	}

	return json.Marshal(m)
}
