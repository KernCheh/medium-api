package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DATABASE_URL string
}

func GetConfig() *Configuration {
	config := &Configuration{}
	gonfig.GetConf(getFileName(false), config)

	gonfig.GetConf(getFileName(true), config)

	return config
}

func getFileName(checkEnv bool) string {
	env := ""
	filename := []string{"config.json"}

	if checkEnv {
		env = os.Getenv("ENV")
		if len(env) == 0 {
			env = "development"
		}

		filename = []string{"config.", env, ".json"}
	}

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
