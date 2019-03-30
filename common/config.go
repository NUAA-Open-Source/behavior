package common

import (
	"github.com/spf13/viper"
	"log"
	"strings"
	"github.com/fsnotify/fsnotify"
)

func SetConfig () error {
	viper.SetConfigName("behavior") // name of config file (without extension)
	viper.AddConfigPath("conf")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BEHAVIOR")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Println("Fatal error config file:", err)
		return err
	}

	return nil
}

func WatchConfig () error {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	return nil
}