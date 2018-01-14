package main

import "fmt"

//ErrorString is a basic implementation of the error interface
type ErrorString struct {
	s string
}

func (e ErrorString) Error() string {
	return fmt.Sprintf("Error: %s", e.s)
}
