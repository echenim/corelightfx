package utilities

import (
	"encoding/json"
	"io"
	"os"
)

func GetConfig() map[string]interface{} {
	j, _ := os.Open("../utilities/config.json")
	defer j.Close()
	b, _ := io.ReadAll(j)
	var settings map[string]interface{}
	json.Unmarshal([]byte(b), &settings)
	return settings
}
