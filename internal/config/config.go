package config

import (
	"encoding/json"

	"github.com/toledompm/http_monitor/pkg/jsonutil"
	"github.com/toledompm/http_monitor/pkg/monitor"
)

func ReadConfig() ([]monitor.MonitorConfig, error) {
	var configs []monitor.MonitorConfig

	jsonByteData, err := jsonutil.OpenJsonFile("configs/config.json")

	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonByteData, &configs)

	return configs, nil
}
