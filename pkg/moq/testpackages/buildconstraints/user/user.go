package user

import "github.com/pehowell/buildconstraints"

// Service does something good with computers.
type Service interface {
	DoSomething(buildconstraints.SomeType) error
}

