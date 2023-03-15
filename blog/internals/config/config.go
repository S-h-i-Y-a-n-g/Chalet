package config

import (
	"chalet/pkg/entity"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

// 整个blog系统的配置
type appConfig struct {
	Mysql entity.MysqlConfig `yaml:"mysql"`
	Redis entity.RedisConfig `yaml:"redis"`
	//Kafka     entity.Kafka       `yaml:"kafka"`
	//Git       entity.Git         `yaml:"git"`
	QQMail entity.QQMail
	// jwt
	JWT     entity.Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}


var AppConfig appConfig

func Init() {
	var env string

	flag.StringVar(&env, "env", "dev", "set environment type : dev, stage, prod")
	flag.Parse()
	// TODO 这里的参数加载，后续可以考虑使用Viper
	InitWithEnv(env)
}

func InitWithEnv(env string) {
	configFileName := getConfigFileByEnv(env)
	configPath := findConfig(configFileName)
	if _, err := os.Stat(configPath); err != nil {
		configPath = findConfigFromCurrentPath(configFileName)
	}

	InitConfig(configPath)
}

func getConfigFileByEnv(env string) string {
	a:="please provide env info to run DLinter. e.g. ./DLinter -env dev"
	b := make([]interface{}, 1, 1)
	b[0] = a
	if env == "none" {
		panic(b[0])
	}
	logFileName := env + ".yml"
	log.Printf("current env is: %s, log file nam is: %s", env, logFileName)
	return logFileName
}

func findConfig(configFileName string) string {
	file, _ := exec.LookPath(os.Args[0])
	log.Println("process name: ", file)

	configPath := path.Join(filepath.Dir(file), "configs", configFileName)
	log.Println("find default config file: ", configPath)
	return configPath
}

func findConfigFromCurrentPath(logFileName string) string {
	_, dir, _, _ := runtime.Caller(0)
	configPath := path.Join(dir, "/../../../configs/", logFileName)
	log.Println("find config file from current path: ", configPath)
	return configPath
}

func InitConfig(path string) {
	filename, _ := filepath.Abs(path)
	log.Println("load config file: ", filename)

	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("read the config file err! %s \n", err.Error())
	}
	err = yaml.Unmarshal(ymlFile, &AppConfig)
	if err != nil {
		log.Fatalf("the config file is not yaml format %s \n", err.Error())
	}
	log.Printf("[AppConfig] %+v \n", AppConfig)
}
