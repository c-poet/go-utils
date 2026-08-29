// Package json provides JSON conversion helpers.
package json

import (
	stdjson "encoding/json"
	"fmt"
)

// ConvertToMap converts value to a map through JSON serialization. It returns
// nil when value cannot be represented as a JSON object.
func ConvertToMap(value any) map[string]any {
	if value == nil {
		return nil
	}

	data, err := stdjson.Marshal(value)
	if err != nil {
		return nil
	}

	var result map[string]any
	if err := stdjson.Unmarshal(data, &result); err != nil {
		return nil
	}
	return result
}

// ConvertToStruct converts origin to target through JSON serialization. Target
// must be a non-nil pointer accepted by encoding/json.
func ConvertToStruct(origin, target any) error {
	if origin == nil || target == nil {
		return fmt.Errorf("origin and target cannot be nil")
	}

	data, err := stdjson.Marshal(origin)
	if err != nil {
		return fmt.Errorf("marshal origin: %w", err)
	}
	if err := stdjson.Unmarshal(data, target); err != nil {
		return fmt.Errorf("unmarshal target: %w", err)
	}
	return nil
}
