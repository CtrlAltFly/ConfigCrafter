package utils

import (
	"gopkg.in/yaml.v3"
	"strings"
)

func PropertiesToYAML(propertiesData []byte) string {
	lines := strings.Split(string(propertiesData), "\n")
	data := convertToYAML(lines)

	output, err := yaml.Marshal(data)
	if err != nil {
		return "Error converting to YAML"
	}
	return string(output)
}

func convertToYAML(properties []string) map[string]interface{} {
	data := make(map[string]interface{})
	for _, line := range properties {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		keys := strings.Split(parts[0], ".")
		value := parts[1]

		current := data
		for i, key := range keys {
			if i == len(keys)-1 {
				current[key] = value
			} else {
				if _, exists := current[key]; !exists {
					current[key] = make(map[string]interface{})
				}
				current = current[key].(map[string]interface{})
			}
		}
	}
	return data
}

func IsYAMLFile(filename string) bool {
	ext := strings.ToLower(filename)
	return strings.HasSuffix(ext, ".yaml") || strings.HasSuffix(ext, ".yml")
}
