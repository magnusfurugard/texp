package printer

import "fmt"

// Raw prints the result statement as it is. Line separated.
func Raw(result []string) error {
	for _, row := range result {
		fmt.Println(row)
	}
	return nil
}
