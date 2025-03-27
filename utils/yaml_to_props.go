package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

func YAMLToProperties(yamlData []byte) string {
	var data map[string]interface{}
	if err := yaml.Unmarshal(yamlData, &data); err != nil {
		return fmt.Sprintf("Error parsing YAML: %v", err)
	}
	properties := convertToProperties(data, "")
	return strings.Join(properties, "\n")
}

func convertToProperties(data map[string]interface{}, prefix string) []string {
	var properties []string
	for key, value := range data {
		newKey := key
		if prefix != "" {
			newKey = prefix + "." + key
		}
		switch v := value.(type) {
		case map[string]interface{}:
			properties = append(properties, convertToProperties(v, newKey)...) // Recursive call
		case []interface{}:
			for i, item := range v {
				properties = append(properties, fmt.Sprintf("%s[%d]=%v", newKey, i, item))
			}
		default:
			properties = append(properties, fmt.Sprintf("%s=%v", newKey, v))
		}
	}
	return properties
}
