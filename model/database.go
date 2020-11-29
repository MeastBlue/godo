package model

// Database type struct
type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Driver   string `mapstructure:"driver"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSL      string `mapstructure:"ssl"`
}
