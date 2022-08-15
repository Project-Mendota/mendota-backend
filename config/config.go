package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigName(c.Name)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config file changed: %s", in.Name)
	})
}

func (c *Config) Init() error {
	err := c.initConfig()
	if err != nil {
		return err
	}
	c.watchConfig()
	return nil
}
