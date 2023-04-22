package hamrobazar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var databasePath = "database.json"

func init() {
	dirname, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	databasePath = fmt.Sprintf("%s/hamrobazar/%s", dirname, databasePath)
}

func readDatabase() ([]string, error) {
	data, err := ioutil.ReadFile(databasePath)
	if err != nil {
		if os.IsNotExist(err) {
			dir := filepath.Dir(databasePath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return nil, fmt.Errorf("os.MkdirAll: %w", err)
			}

			file, err := os.Create(databasePath)
			if err != nil {
				return nil, fmt.Errorf("os.Create: %w", err)
			}
			file.Close()

			return []string{}, nil
		}

		return nil, err
	}

	var content []string
	err = json.Unmarshal(data, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func storeToDatabase(db []string) error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(databasePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
