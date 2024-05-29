package json

import (
	"encoding/json"
	"fmt"
	"io"
)

// NewEncoder delegates to json.NewEncoder
// It is only here so this package can be a drop-in for common encoding/json uses
func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

// Marshal delegates to json.Marshal
// It is only here so this package can be a drop-in for common encoding/json uses
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// limit recursive depth to prevent stack overflow errors
const maxDepth = 10000

// Unmarshal unmarshals the given data.
// Object keys are case-sensitive.
// Numbers decoded into interface{} fields are converted to int64 or float64.
//func Unmarshal(data []byte, v interface{}) error {
//	return kjson.UnmarshalCaseSensitivePreserveInts(data, v)
//}

// ConvertInterfaceNumbers converts any json.Number values to int64 or float64.
// Values which are map[string]interface{} or []interface{} are recursively visited
func ConvertInterfaceNumbers(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = convertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumbers(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

// ConvertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// convertNumber converts a json.Number to an int64 or float64, or returns an error
func convertNumber(n json.Number) (interface{}, error) {
	// Attempt to convert to an int64 first
	if i, err := n.Int64(); err == nil {
		return i, nil
	}
	// Return a float64 (default json.Decode() behavior)
	// An overflow will return an error
	return n.Float64()
}