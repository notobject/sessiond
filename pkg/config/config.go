package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

// Config 全局配置
type Config struct {
	Debug bool        `yaml:"debug"`
	Log   LogConfig   `yaml:"log"`
	Addr  AddrConfig  `yaml:"addr"`
	Redis RedisConfig `yaml:"redis"`
}

// LogConfig 配置
type LogConfig struct {
	Level  string // 日志等级
	Output string // 输出文件
}

// AddrConfig 服务地址
type AddrConfig struct {
	RPC string `yaml:"rpc"`
	Web string `yaml:"web"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

var (
	cfg  *Config
	lock sync.Mutex
)

// GetConfig 获取配置
func GetConfig() *Config {
	return cfg
}

// Reload 加载配置
func Reload(path string) (*Config, error) {
	lock.Lock()
	defer lock.Unlock()

	cfg = &Config{}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return cfg, err
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		return cfg, err
	}

	absPath, err := filepath.Abs(path)
	log.Println("Using config ", absPath)
	return cfg, nil
}
