package test

import (
	"testing"

	"github.com/HrvojeCosic/godb/src"
)

func TestHelloWorld(t *testing.T) {
	if (src.HelloWorld() != "Hello World") {
		t.Error("Expected 'Hello World")
	}
}