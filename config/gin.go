package config

type GinConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (r GinConfig) GetGinAddr() string {
	return r.Host + ":" + r.Port
}
