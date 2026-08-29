package types

import "testing"

func TestIsStruct(t *testing.T) {
	type sample struct{}
	if !IsStruct(sample{}) || !IsStruct(&sample{}) {
		t.Fatal("IsStruct did not recognize a struct")
	}
	if IsStruct(nil) || IsStruct(1) || IsStruct((*sample)(nil)) {
		t.Fatal("IsStruct recognized a non-struct value")
	}
}
