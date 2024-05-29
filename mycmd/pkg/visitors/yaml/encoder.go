package yaml

import "gopkg.in/yaml.v3"

func Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
