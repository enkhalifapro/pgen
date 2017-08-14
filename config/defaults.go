package config

import (
	"github.com/spf13/viper"
)

func init() {
	// Server settings.
	viper.SetDefault("port", 3030)
	viper.SetDefault("path", ".")
	viper.SetDefault("domain", "localhost")

	// MongoDB settings.
	// URI in format 'mongodb://USER:PASSWg@HOST:PORT,HOST:PORT/DBNAME"'.
	viper.SetDefault("db.uri", "mongodb://localhost/TEST")
	viper.SetDefault("db.tls.enable", false)

	// log settings.
	// level can be ERROR|WARNING|INFO|DEBUG.
	viper.SetDefault("log.level", "ERROR")
	viper.SetDefault("log.ginrus", true)

	// SimilarTech settings.
	viper.SetDefault("similar_tech.key", "apiKey")
}
