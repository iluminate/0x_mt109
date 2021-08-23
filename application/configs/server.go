package configs

import "0x_mt109/helpers/loader"

type ConfigServer struct {
	Server struct {
		Port        int32  `yaml:"port"`
		ContextPath string `yaml:"contextPath"`
	} `yaml:"server"`
}

func (conf *ConfigServer) Merge(envCfg interface{}) loader.ConfigurationProperties {
	envConfig := envCfg.(*ConfigServer)
	conf.Server.Port = loader.GetVal(envConfig.Server.Port, conf.Server.Port).(int32)
	conf.Server.ContextPath = loader.GetVal(envConfig.Server.ContextPath, conf.Server.ContextPath).(string)
	return conf
}
