package priority

import "fmt"

type Priority int

const (
	_ Priority = iota
	LOW
	MIDDLE
	HIGH
)

func New(ip int32) (Priority, error) {
	p := Priority(ip)
	if !p.isDefined() {
		return -1, fmt.Errorf("%d is undefined at Priority", ip)
	}
	return p, nil
}

func (p Priority) isDefined() bool {
	switch p {
	case LOW, MIDDLE, HIGH:
		return true
	default:
		return false
	}
}
