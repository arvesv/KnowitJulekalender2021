package main

import "testing"

func TestGetNumberEn(t *testing.T) {
	s, l := getnumber("en", 0)
	if s != 1 {
		t.Fatalf("Expected 1. Got %v", s)
	}
	if l != 2 {
		t.Fatalf("Length of en should be 2. Got %v", l)
	}
}

func TestGetNumberTre(t *testing.T) {
	s, l := getnumber("tre", 0)
	if s != 3 {
		t.Fatalf("Expected 3. Got %v", s)
	}
	if l != 3 {
		t.Fatalf("Length of en should be 3. Got %v", l)
	}
}

func TestGetMulti(t *testing.T) {
	s, l := getnumber("treen", 0)
	if s != 3 {
		t.Fatalf("Expected 3. Got %v", s)
	}
	if l != 3 {
		t.Fatalf("Length of en should be 3. Got %v", l)
	}
}
