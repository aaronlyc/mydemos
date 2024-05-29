package unstructured

import (
	"mydemos/mycmd/pkg/visitors/json"
	"mydemos/mycmd/pkg/visitors/yaml"
)

type Unstructured struct {
	// Object is a JSON compatible map with string, float, int, bool, []interface{}, or
	// map[string]interface{}
	// children.
	Object map[string]interface{}
}

// MarshalJSON ensures that the unstructured object produces proper
// JSON when passed to Go's standard JSON library.
func (u *Unstructured) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Object)
}

func (u *Unstructured) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(u.Object)
}

//// UnmarshalJSON ensures that the unstructured object properly decodes
//// JSON when passed to Go's standard JSON library.
//func (u *Unstructured) UnmarshalJSON(b []byte) error {
//	_, _, err := UnstructuredJSONScheme.Decode(b, nil, u)
//	return err
//}
