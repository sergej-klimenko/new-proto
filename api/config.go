package api

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

type Settings struct {
	MongoHost string `mapstructure:"MONGO_HOST"`
	MongoName string `mapstructure:"MONGO_NAME"`
	MongoUser string `mapstructure:"MONGO_USER"`
	MongoPass string `mapstructure:"MONGO_PASS"`
}

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

	return nil
}

func GetConfig(key string) string {
	return fmt.Sprintf("%v", viper.Get(key))
}
