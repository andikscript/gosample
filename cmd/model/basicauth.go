package model

var AuthBasic *Basic

type Basic struct {
	BasicAuth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"basicauth"`
}

func GetBasicAuth() *Basic {
	return AuthBasic
}
