package service

import (
	"testing"
)

func TestGetTree(t *testing.T) {
	_, err := NewHomeService().GetTree()
	if err != nil {
		t.Error(err)
	}
}
