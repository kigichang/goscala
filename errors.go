package goscala

import "fmt"

var (
	ErrUnsupported = fmt.Errorf("unsupported")
	ErrUnsatisfied = fmt.Errorf("unsatisfied")
	ErrEmpty = fmt.Errorf("emtpy")
	ErrLeft = fmt.Errorf("left")
)