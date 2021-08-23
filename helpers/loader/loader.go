package loader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

const configDir = "config"
const prefixConfigFile = "conf"

var configProps = make(map[string]ConfigurationProperties)

func Load(cfgs ...ConfigurationProperties) {
	configFiles := loadConfigs(configProps, cfgs...)
	fmt.Printf("Loaded configs file: [%s]\n", strings.Join(*configFiles, ","))
}

func ReadConf(conf ConfigurationProperties) {
	typeName := reflect.TypeOf(conf).Elem().Name()
	config := configProps[typeName]
	if config == nil {
		fmt.Printf("Load Config:[%s]\n", typeName)
		configFiles := loadConfigs(configProps, conf)
		config = configProps[typeName]
		fmt.Printf("Loaded config file: [%s]\n", strings.Join(*configFiles, ","))
	} else {
		fmt.Printf("Read Config:[%s]\n", typeName)
	}
	src := reflect.ValueOf(config).Elem()
	dst := reflect.ValueOf(conf).Elem()
	dst.Set(src)

}

func loadConfigs(configProps map[string]ConfigurationProperties, cfgs ...ConfigurationProperties) *[]string {
	configNames := make([]string, 0)
	baseConfFileName := getConfigFileNames()
	for _, cfg := range cfgs {
		unmarshalConf(baseConfFileName, cfg)
		typeName := reflect.TypeOf(cfg).Elem().Name()
		configProps[typeName] = cfg
	}
	configNames = append(configNames, baseConfFileName)
	return &configNames
}

func unmarshalConf(configFile string, configuration ConfigurationProperties) {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("error reading yaml file   #%v ", err)
	}
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))
	err = yaml.Unmarshal(yamlFile, configuration)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

}

func getConfigFileNames() string {
	return fmt.Sprintf("%s/%s.yaml", configDir, prefixConfigFile)
}

func GetVal(value interface{}, defaultValue interface{}) interface{} {
	if value == nil || value == "" || value == 0 || value == false {
		return defaultValue
	}
	return value
}

type ConfigurationProperties interface {
	Merge(envCfg interface{}) ConfigurationProperties
}