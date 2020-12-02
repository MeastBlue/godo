package model

// TLS type struct
type TLS struct {
	Path string `mapstructure:"path"`
	Crt  string `mapstructure:"crt"`
	Key  string `mapstructure:"key"`
}
