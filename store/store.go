package store

import (
	"os"

	"github.com/spf13/viper"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func WriteBookmark(config string, bookmark int) error {
	if fileExists(config) == false {
		os.Stderr.WriteString("Warning: No config file found to persist bookmark. Create it with `init`\n")
		return nil
	}

	viper.GetViper().Set("bookmark", bookmark)
	err := viper.WriteConfig()

	if err != nil {
		return err
	}

	return nil
}

func GetBookmark() int {
	bookmark := viper.GetInt("bookmark")
	return bookmark
}
