package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {

	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)

}

func TestViperJSON(t *testing.T) {

	var config *viper.Viper = viper.New()
	
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "farhan", config.GetString("app.author"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))

}

func TestViperYMAL(t *testing.T) {

	var config *viper.Viper = viper.New()
	
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "farhan", config.GetString("app.author"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))

}

func TestViperENV(t *testing.T) {

	var config *viper.Viper = viper.New()
	
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "farhan", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))

}