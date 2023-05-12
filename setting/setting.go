package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release bool `ini:"release"`
	Port    int  `ini:"port"`
}

// Init
//
//	@Description: 初始化配置
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
