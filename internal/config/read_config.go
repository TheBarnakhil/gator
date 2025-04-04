package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error while fetching file: %w", err)
	}

	var data Config
	if err := json.Unmarshal(file, &data); err != nil {
		return Config{}, fmt.Errorf("error while unmarshaling file: %w", err)
	}
	return data, nil
}
