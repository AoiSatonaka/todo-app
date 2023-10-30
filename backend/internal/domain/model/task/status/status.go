package status

import "fmt"

type Status int32

const (
	_ Status = iota
	TODO
	DOING
	DONE
)

func New(is int32) (Status, error) {
	s := Status(is)
	if !s.isDefined() {
		return -1, fmt.Errorf("%d is undefined at Status", is)
	}
	return s, nil
}

func (s Status) isDefined() bool {
	switch s {
	case TODO, DOING, DONE:
		return true
	default:
		return false
	}
}
