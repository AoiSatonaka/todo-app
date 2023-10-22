package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if Hello("world") != "hello world" {
		t.Error("Expected Hello to append 'world'")
	}
}
