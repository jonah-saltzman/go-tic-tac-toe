package main

import "testing"

func TestMaxInt(t *testing.T) {
	m := MaxInt(-50, 500, 10000, -10000000)
	if m != 10000 {
		t.Errorf("MaxInt failed")
	}
	m = MaxInt(-29385, -1, -50)
	if m != -1 {
		t.Errorf("MaxInt failed")
	}
	m = MaxInt(10)
	if m != 10 {
		t.Errorf("MaxInt failed")
	}
}

func TestMinInt(t *testing.T) {
	m := MinInt(-50, 500, 10000, -10000000)
	if m != -10000000 {
		t.Errorf("MinInt failed")
	}
	m = MinInt(-29385, -1, -50)
	if m != -29385 {
		t.Errorf("MinInt failed")
	}
	m = MinInt(10)
	if m != 10 {
		t.Errorf("MinInt failed")
	}
}
