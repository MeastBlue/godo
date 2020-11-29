package model

// JWT type struct
type JWT struct {
	Access  string `mapstructure:"access"`
	Refresh string `mapstructure:"refresh"`
}

// Token type struct
type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails type struct
type AccessDetails struct {
	AccessUuid string
	UserId     string
}
