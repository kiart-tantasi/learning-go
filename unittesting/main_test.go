package unittesting

import "testing"

// should pass
func TestAddPass(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Error("FAILED TestAddPass")
	}
}

// this test will not be run because function name is not prefixed by "Test"
func AddNotRun(t *testing.T) {
	if Add(1, 1) != -999 {
		t.Error("FAILED AddNotRun")
	}
}

// should fail
func TestAddFail(t *testing.T) {
	if Add(1, 1) != -999 {
		t.Error("FAILED TestAddFail")
	}
}
