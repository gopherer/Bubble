package tools

import "gopkg.in/ini.v1"

var Cfg = new(Config)

// Config 应用程序配置
type Config struct {
	Release  bool         `ini:"release" json:"release"`
	Port     int          `ini:"port" json:"port"`
	MysqlCfg *MysqlConfig `ini:"mysql" json:"mysql"`
}

//MysqlConfig 数据库配置
type MysqlConfig struct {
	User    string `ini:"user" json:"user"`
	Pwd     string `ini:"pwd" json:"pwd"`
	DbName  string `ini:"db_name" json:"db_name"`
	Host    string `ini:"host" json:"host"`
	Port    int    `ini:"port" json:"port"`
	Charset string `ini:"charset" json:"charset"`
}

func Init(file string) error {
	return ini.MapTo(Cfg, file)
}
