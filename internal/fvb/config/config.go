package config

type Config struct {
	MongoDB MongoCfg `json:"mongo"`
	Bot     Bot      `json:"bot"`
}

type Bot struct {
	Token string `json:"token" env:"FVB_TOKEN" required:"true"`
}

type MongoCfg struct {
	Addr     string `json:"addr"`
	HostName string `json:"host_name"`
	Port     string `json:"host_name"`
	DBName   string `json:"db_name"`
}
