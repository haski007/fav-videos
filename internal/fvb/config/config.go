package config

type Config struct {
	MongoDB MongoCfg `json:"mongo"`
	Bot     Bot      `json:"bot"`
}

type Bot struct {
	Token     Token `json:"token" env:"FVB_TOKEN,required"`
	CreatorID int64 `json:"creator_id" env:"CREATOR_ID"`
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
