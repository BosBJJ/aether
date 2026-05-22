package utils

import (
	"aether/cmd/config"
	"encoding/json"
	"fmt"
)


func PrintJSON(v any) error {
	var data []byte
	var err error
	if config.Pretty {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}
	fmt.Println(string(data))
    return nil
}