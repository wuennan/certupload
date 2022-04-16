package conf_test

import (
	"fmt"
	"github.com/wuennan/certupload/conf"
	"testing"
)

func Test_load(t *testing.T) {
	conf.LoadConfigFromToml("../etc/config.toml")
	c := conf.C()
	fmt.Println(c.RegionId)
}
