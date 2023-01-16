package golearn

import "testing"

func TestAbsolute(t *testing.T) {
	c := Absolute(-5)
	if c != 5 {
		t.Logf("expect 5, but got %d", c)
		t.Fail()
	}
}

func TestAbsolute_WithPositive(t *testing.T) {
	c := Absolute(5)
	if c != 5 {
		t.Logf("expect 5, but got %d", c)
		t.Fail()
	}
}
