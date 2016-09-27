package main

import "testing"

func TestAdd(t *testing.T) {
	if add(7, 11) != 18 {
		t.Fail()
	}
	if add(256, 276) != 532 {
		t.Fail()
	}
	if add(17985, 15568) != 33553 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	if sub(7, 11) != -4 {
		t.Fail()
	}
	if sub(256, 276) != -20 {
		t.Fail()
	}
	if sub(17985, 15568) != 2417 {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	if mul(7, 11) != 77 {
		t.Fail()
	}
	if mul(256, 276) != 70656 {
		t.Fail()
	}
	if mul(17985, 15568) != 279990480 {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	if div(5, 2) != 2 {
		t.Fail()
	}
	if div(250, 50) != 5 {
		t.Fail()
	}
	if div(3, 5) != 0 {
		t.Fail()
	}
}
