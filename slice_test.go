package go_helpers

import (
	"testing"
)

func TestMakeBoolSliceUnqiue(t *testing.T) {
	r := []bool{true, false}
	a := []bool{true, false, false, true}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
}

func TestMakeIntSliceUnique(t *testing.T) {
	r := []int{0, 1, 2, 3}
	a := []int{0, 1, 2, 3, 3, 2, 1}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
}

func TestMakeFloat32SliceUnique(t *testing.T) {
	r := []float32{0.01, 1.22, 0.0}
	a := []float32{0, 0.0, 1.22, 0.01, 0.01, 1.22, 0.0, 0}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
}

func TestMakeFloat64SliceUnique(t *testing.T) {
	r := []float64{0.001, 1.2222, 0.000, 0.01}
	a := []float64{0, 0.00, 1.2222, 0.01, 0.001, 1.2222, 0.0, 0}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
}

func TestMakeStringSliceUnqiue(t *testing.T) {
	r := []string{"a", "b", "c"}
	a := []string{"a", "b", "c", "a", "b", "c", "c"}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
}
