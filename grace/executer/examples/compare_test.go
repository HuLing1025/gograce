package executer

import "testing"

func TestCompare(t *testing.T) {
	if Compare(1, 2) != -1 {
		t.Error("Demo(1, 2) != -1")
	}
	if Compare(1, 1) != 0 {
		t.Error("Demo(1, 1) != 0")
	}
}
