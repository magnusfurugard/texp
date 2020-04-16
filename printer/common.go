package printer

// TODO: Tests

import "fmt"

// ToMapInterface takes incomming arguments and adds them tp a map[string]interface{}.
func ToMapInterface(tokens map[string][]string, result []string) map[string]interface{} {
	m := map[string]interface{}{}
	m["result"] = result
	m["tokens"] = tokens
	return m
}

// Print prints the result, and tokens depending on the selected method.
func Print(method string, tokens map[string][]string, result []string) error {
	m := ToMapInterface(tokens, result)
	switch method {
	case "raw":
		Raw(result)
	case "json":
		JSON(m)
	case "yaml":
		YAML(m)
	default:
		return fmt.Errorf("invalid output format `%v`. Must be either `raw`, `json` or `yaml`", method)
	}

	return nil
}
