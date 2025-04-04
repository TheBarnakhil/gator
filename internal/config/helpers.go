package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error while fetching home path: %w", err)
	}
	return filepath.Join(home, configFileName), nil
}

func writeToFile(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error while marshalling json: %w", err)
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0666)
	return err
}
