package visitors

import (
	"bytes"
	"os"
	"testing"
)

func TestStreamVisitor(t *testing.T) {
	file, err := os.ReadFile("./demo.yaml")
	if err != nil {
		t.Errorf(err.Error())
	}

	visitor := NewStreamVisitor(bytes.NewReader(file), "./demo.yaml")
	err = visitor.Visit(nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
