package os

import "testing"

func TestGetProgramPath(t *testing.T) {
	if path := GetProgramPath(); path == "" {
		t.Fatal("GetProgramPath returned an empty path")
	}
}

func TestGetProgramExtPathMissing(t *testing.T) {
	if _, err := GetProgramExtPath("go-utils-test-program-that-does-not-exist"); err == nil {
		t.Fatal("GetProgramExtPath accepted a missing executable")
	}
}
