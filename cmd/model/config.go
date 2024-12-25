package model

var Shared *Config

type Config struct {
	Server struct {
		Port         int `json:"port"`
		ReadTimeout  int `json:"read_timeout"`
		WriteTimeout int `json:"write_timeout"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
		Schema   string `json:"schema"`
	}
}

func GetConfig() *Config {
	return Shared
}
