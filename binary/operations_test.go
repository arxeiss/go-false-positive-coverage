package binary_test

import (
	"testing"

	. "github.com/arxeiss/go-false-positive-coverage/binary"
)

func TestAdd(t *testing.T) {
	if res := Add(2.5, 4.8); res != 7.3 {
		t.Errorf("Expected result to be 7.3, got %f", res)
	}
}

func TestSub(t *testing.T) {
	if res := Sub(2.5, 4.8); res != -2.3 {
		t.Errorf("Expected result to be 2.3, got %f", res)
	}
}

func TestMul(t *testing.T) {
	if res := Mul(4.5, 8); res != 36 {
		t.Errorf("Expected result to be 36, got %f", res)
	}
}

func TestDiv(t *testing.T) {
	if res := Div(12.5, 2.5); res != 5 {
		t.Errorf("Expected result to be 5, got %f", res)
	}
}
