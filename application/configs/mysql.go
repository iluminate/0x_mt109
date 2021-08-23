package configs

import (
	"0x_mt109/helpers/loader"
	"time"
)

type MysqlConfig struct {
	Mysql struct {
		Host            string
		Port            string
		User            string
		Password        string
		Database        string
		ConnMaxLifetime	time.Duration
		MaxIdleConns   	int
		MaxOpenConns   	int
	}
}

func (conf *MysqlConfig) Merge(envCfg interface{}) loader.ConfigurationProperties {
	envConfig := envCfg.(*MysqlConfig)
	conf.Mysql.Host = loader.GetVal(envConfig.Mysql.Host, conf.Mysql.Host).(string)
	conf.Mysql.Port = loader.GetVal(envConfig.Mysql.Port, conf.Mysql.Port).(string)
	conf.Mysql.User = loader.GetVal(envConfig.Mysql.User, conf.Mysql.User).(string)
	conf.Mysql.Password = loader.GetVal(envConfig.Mysql.Password, conf.Mysql.Password).(string)
	conf.Mysql.Database = loader.GetVal(envConfig.Mysql.Database, conf.Mysql.Database).(string)
	conf.Mysql.ConnMaxLifetime = loader.GetVal(envConfig.Mysql.ConnMaxLifetime, conf.Mysql.ConnMaxLifetime).(time.Duration)
	conf.Mysql.MaxIdleConns = loader.GetVal(envConfig.Mysql.MaxIdleConns, conf.Mysql.MaxIdleConns).(int)
	conf.Mysql.MaxOpenConns = loader.GetVal(envConfig.Mysql.MaxOpenConns, conf.Mysql.MaxOpenConns).(int)
	return conf
}