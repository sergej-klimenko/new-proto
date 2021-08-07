package api

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

type Settings struct {
	Environment string `mapstructure:"Environment"`
}

var ErrInvalidConfiguration = errors.New("invalid configuration")

func LoadConfiguration() error {
	s := &Settings{}

	viper.SetConfigFile("settings.json")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println("error reading config file, using environment variables.")
	}

	if err := viper.Unmarshal(&s); err != nil {
		return errors.Wrap(err, "error unmarshaling settings file")
	}

	if s.Environment == "" {
		return ErrInvalidConfiguration
	}

	return nil
}

func GetConfig(key string) string {
	return fmt.Sprintf("%v", viper.Get(key))
}
