package model

// JWT type struct
type JWT struct {
	Access  string `mapstructure:"access"`
	Refresh string `mapstructure:"refresh"`
}
