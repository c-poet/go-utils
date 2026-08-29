package strings

import (
	"reflect"
	"testing"
)

func TestToArrayNonEmpty(t *testing.T) {
	got := ToArrayNonEmpty(",one,,two,", ",")
	want := []string{"one", "two"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ToArrayNonEmpty() = %v, want %v", got, want)
	}
}

func TestToArrayComma(t *testing.T) {
	got := ToArrayComma("one,,two")
	want := []string{"one", "", "two"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ToArrayComma() = %v, want %v", got, want)
	}
}
