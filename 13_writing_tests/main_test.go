package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expected-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Did not get an error when we should have")
			}
		} else {
			if err != nil {
				t.Error("Got an error when we should not have", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

//	//	//	//	//

//	test manual

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0.0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }
