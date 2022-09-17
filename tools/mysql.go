package tools

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *Config) (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		cfg.MysqlCfg.User, cfg.MysqlCfg.Pwd, cfg.MysqlCfg.Host, cfg.MysqlCfg.Port, cfg.MysqlCfg.DbName, cfg.MysqlCfg.Charset)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
