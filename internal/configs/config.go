package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type resource struct {
	Name     string
	Endpoint string
	DestUrl  string
}

type configuration struct {
	Server struct {
		Host       string
		ListenPort string
	}
	Resources []resource
}

func NewConfig() (config *configuration, err error) {
	viper.AddConfigPath("settings")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading config file: %s", err)
	}
	if err = viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	return
}
