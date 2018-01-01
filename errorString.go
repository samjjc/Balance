package main

import "fmt"

type ErrorString struct {
	S string
}

func (e ErrorString) Error() string {
	return fmt.Sprintf("work failed: %s", e.S)
}
