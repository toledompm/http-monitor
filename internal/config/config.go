package config

import (
	"encoding/json"

	"github.com/toledompm/http_monitor/internal/monitor"
	"github.com/toledompm/http_monitor/pkg/jsonutil"
)

func ReadConfig(configFilePath string) ([]monitor.MonitorConfig, error) {
	var configs []monitor.MonitorConfig

	jsonByteData, err := jsonutil.OpenJsonFile(configFilePath)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonByteData, &configs)

	return configs, nil
}
