// configuration handling
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Description string   `json:"description"`
	Command     string   `json:"command"`
	DependsOn   []string `json:"dependsOn"`
}

type Config struct {
	Tasks map[string]Task `json:"tasks"`
}

const ConfigFileName = "taskifile.json"

func Load() (*Config, error) {
	file, err := os.ReadFile(ConfigFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error: the config file '%s' doesn't exist.", ConfigFileName)
		}
		if os.IsPermission(err) {
			return nil, fmt.Errorf("error: we don't have permission to read the config file '%s'.", ConfigFileName)
		}
		return nil, fmt.Errorf("error: failed to read the config file %s: %w", ConfigFileName, err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		// TODO: handle all other errors like invalid json file, no content, etc
		return nil, fmt.Errorf("failed to parse the config file '%s'. Please make sure it has a valid json", ConfigFileName)
	}

	return &config, nil
}

func (c *Config) GetTask(name string) (Task, bool) {
	task, exists := c.Tasks[name]
	return task, exists
}

func (c *Config) ListTaskNames() []string {
	names := make([]string, 0, len(c.Tasks))
	for name := range c.Tasks {
		names = append(names, name)
	}

	return names
}
