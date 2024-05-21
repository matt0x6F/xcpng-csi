package config

import (
	"os"

	"github.com/matt0x6f/xcpng-csi/internal/structs"
	"gopkg.in/yaml.v2"
)

const configLocation = "/config/xcpng-csi.conf"

// Load loads the XCP-ng config
func Load() (*structs.Config, error) {
	config := structs.Config{
		NodeID:    os.Getenv("NODE_ID"),
		ClusterID: os.Getenv("CLUSTER_ID"),
	}

	yamlFile, err := os.ReadFile(configLocation)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
