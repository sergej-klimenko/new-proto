package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type Settings struct {
	Environment string ` mapstructure:"ENVIRONMENT"`
	list        map[string]string
}

var settings = &Settings{
	list: map[string]string{
		"Environment": "",
	},
}

var ErrInvalidConfiguration = errors.New("invalid configuration")

func Load() error {
	settingsFile, _ := ioutil.ReadFile("settings.json")
	if err := json.Unmarshal(settingsFile, settings); err != nil {
		return errors.Wrap(err, "error reading configuration file.")
	}

	if os.Getenv("ENVIRONMENT") != "" {
		settings.Environment = os.Getenv("ENVIRONMENT")
	}

	settings.list["ENVIRONMENT"] = settings.Environment

	if settings.Environment == "" {
		return ErrInvalidConfiguration
	}

	return nil
}

func Get(key string) string {
	return settings.list[key]
}
