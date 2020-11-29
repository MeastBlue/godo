package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/meastblue/godo/model"
	"github.com/spf13/viper"
)

// Init conf function
func Init(ext, file, path string) {
	v := viper.New()
	c := model.Config{}

	v.SetConfigType(ext)
	v.SetConfigName(file)
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't load config: %s", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("Couldn't read config: %s", err)
	}

	setServerConf(&c)
	setDatabaseConf(&c)
	setStorageConf(&c)
	setJwtConf(&c)
}

// setServerConf function
func setServerConf(c *model.Config) {
	v := reflect.ValueOf(c.Server)
	field := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := fmt.Sprintf("srv.%s", field.Field(i).Name)
		val := fmt.Sprintf("%s", v.Field(i).Interface())
		os.Setenv(key, val)
	}
}

// setDatabaseConf function
func setDatabaseConf(c *model.Config) {
	v := reflect.ValueOf(c.Db)
	field := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := fmt.Sprintf("db.%s", field.Field(i).Name)
		val := fmt.Sprintf("%s", v.Field(i).Interface())
		os.Setenv(key, val)
	}
}

func setStorageConf(c *model.Config) {
	v := reflect.ValueOf(c.Storage)
	field := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := fmt.Sprintf("stg.%s", field.Field(i).Name)
		val := fmt.Sprintf("%s", v.Field(i).Interface())
		os.Setenv(key, val)
	}
}

func setJwtConf(c *model.Config) {
	v := reflect.ValueOf(c.JWT)
	field := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := fmt.Sprintf("jwt.%s", field.Field(i).Name)
		val := fmt.Sprintf("%s", v.Field(i).Interface())
		os.Setenv(key, val)
	}
}
