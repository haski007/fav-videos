package config

type Config struct {
	MongoDB MongoCfg `json:"mongo"`
	Bot     Bot      `json:"bot"`
	TikTok  TikTok   `json:"TikTok"`
}

type Bot struct {
	Token     Token `json:"token" env:"FVB_TOKEN,required"`
	CreatorID int64 `json:"creator_id" env:"CREATOR_ID"`
}

type TikTok struct {
	// Deprecated
	ChannelID int64 `env:"CHANNEL_ID"`

	Username  string `env:"TIKTOK_USERNAME,required"`
	SecUserID string `env:"TIKTOK_SEC_USER_ID"`
}

type MongoCfg struct {
	Addr     string `json:"addr" env:"MONGO_ADDR"`
	HostName string `json:"host_name" env:"MONGO_HOST"`
	Port     string `json:"port" env:"MONGO_PORT"`
	DBName   string `json:"db_name" env:"MONGO_DBNAME"`
	Username string `json:"username" env:"MONGO_USERNAME"`
	Password string `json:"password" env:"MONGO_PASSWORD"`
}

func (b Bot) GetToken() Token {
	return b.Token
}

type Token string

func (t Token) String() string {
	return string(t)
}
