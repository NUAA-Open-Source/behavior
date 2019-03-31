package common

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	logFile *os.File
	err     error
)

func InitLogger() {

	os.Mkdir("log", os.ModePerm|os.ModeDir)
	logFile, err = os.OpenFile("log/behavior.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if !viper.GetBool("basic.debug") {
		log.SetOutput(logFile)
	}
}

func GetLogFile() *os.File {
	return logFile
}
