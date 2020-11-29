package model

// Config type struct
type Config struct {
	Server  Server   `mapstructure:"server"`
	JWT     JWT      `mapstructure:"jwt"`
	Storage Storage  `mapstructure:"storage"`
	Db      Database `mapstructure:"database"`
}
