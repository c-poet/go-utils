package json

import "testing"

type conversionSource struct {
	Name string `json:"name"`
}

func TestConversions(t *testing.T) {
	source := conversionSource{Name: "go-utils"}
	if got := ConvertToMap(source)["name"]; got != "go-utils" {
		t.Fatalf("map value = %v, want go-utils", got)
	}

	var target conversionSource
	if err := ConvertToStruct(map[string]string{"name": "converted"}, &target); err != nil {
		t.Fatalf("ConvertToStruct: %v", err)
	}
	if target.Name != "converted" {
		t.Fatalf("target.Name = %q, want converted", target.Name)
	}
}
