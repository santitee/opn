package assignment2_test

import (
	"opn/assignment2"
	"testing"
)

func TestIsValidTrue(t *testing.T) {
	v := assignment2.Is_valid_ip("172.16.254.1")
	if true != v {
		t.Errorf("result is %v", v)
	}
}

func TestIsValidFalse(t *testing.T) {
	v := assignment2.Is_valid_ip("172.316.254.1")
	if false != v {
		t.Errorf("result is %v", v)
	}
}
