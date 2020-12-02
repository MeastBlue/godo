package model

// Config type struct
type Config struct {
	Server  Server   `mapstructure:"server"`
	TLS     TLS      `mapstructure:"tls"`
	JWT     JWT      `mapstructure:"jwt"`
	Storage Storage  `mapstructure:"storage"`
	Db      Database `mapstructure:"database"`
}
