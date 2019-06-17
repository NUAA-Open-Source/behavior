package common

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/raven-go"
	"github.com/spf13/viper"
)

func SetConfig() error {
	viper.SetConfigName("behavior") // name of config file (without extension)
	viper.AddConfigPath("conf")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BEHAVIOR")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Println("Fatal error config file:", err)
		raven.CaptureError(err, map[string]string{"type": "config"})
		return err
	}

	return nil
}

func WatchConfig() error {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	return nil
}

func DefaultConfig() error {
	// basic default values
	viper.SetDefault("basic.debug", true)
	viper.SetDefault("basic.maintenance", false)
	viper.SetDefault("basic.port", "8080")
	// csrf default values
	viper.SetDefault("csrf.cookie_secret", "your_cookie_secret")
	viper.SetDefault("session_name", "your_session_name")
	viper.SetDefault("secret", "your_csrf_secret")
	// storage default values
	viper.SetDefault("storage.mysql.user", "root")
	viper.SetDefault("storage.mysql.password", "")
	viper.SetDefault("storage.mysql.host", "localhost")
	viper.SetDefault("storage.mysql.port", "3306")
	viper.SetDefault("sotrage.mysql.database", "behavior")
	viper.SetDefault("storage.mysql.timezone", "Asia%2FShanghai")
	viper.SetDefault("storage.mysql.retry_interval", 20)
	viper.SetDefault("storage.mysql.max_idle_conns", 100)
	viper.SetDefault("storage.mysql.max_open_conns", 100)
	viper.SetDefault("storage.mysql.conn_max_lifetime", 30)
	// sentry default values
	viper.SetDefault("sentry.dsn", "")
	viper.SetDefault("sentry.default_logger_name", "behavior")
	viper.SetDefault("sentry.sample_rate", 1)

	return nil
}
