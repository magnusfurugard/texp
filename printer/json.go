package printer

import (
	"encoding/json"
	"fmt"
)

// JSON marshals the incoming params into json and prints it.
func JSON(m map[string]interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
