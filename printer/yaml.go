package printer

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// YAML prints the tokens and result set as yaml.
func YAML(m map[string]interface{}) error {
	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
