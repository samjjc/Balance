package main

import "fmt"

type ErrorString struct {
	s string
}

func (e ErrorString) Error() string {
	return fmt.Sprintf("Work Failed: %s", e.s)
}
