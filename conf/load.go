package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/wuennan/certupload/app/aliyun/slb"
)

var (
	config *slb.Config
)

func C() *slb.Config {
	return config
}

func LoadConfigFromToml(filePath string) error {
	cfg := slb.NewConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}
	config = cfg
	return nil
}
