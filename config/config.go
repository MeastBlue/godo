package config

import (
	"fmt"
	"io/ioutil"
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
	setTlsConf(&c)
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

func setTlsConf(c *model.Config) {
	v := reflect.ValueOf(c.TLS)
	field := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := fmt.Sprintf("tls.%s", field.Field(i).Name)
		val := fmt.Sprintf("%s", v.Field(i).Interface())
		os.Setenv(key, val)
	}
}

func setJwtConf(c *model.Config) {
	v := reflect.ValueOf(c.JWT)
	field := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := fmt.Sprintf("jwt.%s", field.Field(i).Name)
		path := fmt.Sprintf("%s", v.Field(i).Interface())
		val, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("Couldn't read jwt path: %s", err)
		}
		os.Setenv(key, string(val))
	}
}
