package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteConfig writes json config file, filename.json, to path. Expected input is
// a ConfigFile struct
func WriteConfig(path, filename string, config *ConfigFile) error {
	file := path + "/" + filename + ".json"
	fmt.Println("Writing config file: " + file)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	json, _ := json.Marshal(config)
	f.Write(json)
	f.Sync()

	return nil
}
