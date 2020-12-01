package config

type Config struct {
	MongoDB MongoCfg `json:"mongo"`
	Bot     Bot      `json:"bot"`
	TikTok  TikTok   `json"TikTok"`
}

type Bot struct {
	Token     Token `json:"token" env:"FVB_TOKEN,required"`
	CreatorID int64 `json:"creator_id" env:"CREATOR_ID"`
}

type TikTok struct {
	// TODO: create collection in mongo with all chat ids
	// Deprecated
	ChannelID int64 `env:"CHANNEL_ID,required"`

	Username  string `env:"TIKTOK_USERNAME,required"`
	SecUserID string `env:"TIKTOK_SEC_USER_ID"`
}

type MongoCfg struct {
	Addr     string `json:"addr"`
	HostName string `json:"host_name"`
	Port     string `json:"host_name"`
	DBName   string `json:"db_name"`
}

func (b Bot) GetToken() Token {
	return b.Token
}

type Token string

func (t Token) String() string {
	return string(t)
}
