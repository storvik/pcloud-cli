package config

type ConfigFile struct {
	UserID      int    `json:"userid"`
	AccessToken string `json:"access_token"`
}
