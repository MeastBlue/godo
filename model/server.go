package model

// Server type struct
type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
