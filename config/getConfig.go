package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config 定义配置文件结构体
type Config struct {
	Environment struct {
		Test string `yaml:"test"`
		Pre  string `yaml:"pre"`
		Pro  string `yaml:"pro"`
	} `yaml:"environment"`

	Handle struct {
		Test string `yaml:"test"`
		Pre  string `yaml:"pre"`
		Pro  string `yaml:"pro"`
	} `yaml:"handle"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	Header struct {
		Cookie string `yaml:"cookie"`
	} `yaml:"header"`

	FilePath struct {
		DiscountCodePath string `yaml:"discountCodePath"`
		PreCasePath      string `yaml:"preCasePath"`
	} `yaml:"filePath"`
}

// yaml文件路径
const configPath = "./config/config.yaml"

func GetConfig() Config {
	// 读取 YAML 配置文件
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	// 解析 YAML 配置文件
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

//func main() {
//	config := GetConfig()
//	// 获取yaml文件数据
//	envTest := config.Environment.Test
//	jsonPath := config.FilePath.JsonPath
//	cookie := config.Header.Cookie
//	fmt.Println(envTest)
//	fmt.Println(jsonPath)
//	fmt.Println(cookie)
//}
