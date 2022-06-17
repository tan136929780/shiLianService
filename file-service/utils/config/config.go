package config

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		path, _ := os.Getwd()
		viper.AddConfigPath(path + "/config")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	//初始化运行配置环境
	os.Setenv("VFILE_ENV", viper.GetString("run_mode"))

	//加载环境变量
	viper.SetEnvPrefix("vFile")
	viper.BindEnv("env")
	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}

//返回当前环境名称
func GetEnv() string {
	var env string
	switch viper.GetString("env") {
	case "dev":
		fallthrough
	case "feature":
		fallthrough
	case "preprod":
		fallthrough
	case "prod":
		env = viper.GetString("env")
	default:
		env = "dev"
	}
	//fmt.Println("env=", viper.GetString("env"), "env=", env)
	return env
}

//根据当前环境获取配置文件中的字符串
func GetString(s string) string {
	//fmt.Println(GetEnv()+"."+s)
	r := viper.GetString(GetEnv() + "." + s)
	if r == "" {
		r = viper.GetString("common." + s)
	}
	return r
}

//根据当前环境获取配置文件中的整形
func GetInt(s string) int {
	//fmt.Println(GetEnv()+"."+s)
	r := viper.GetInt(GetEnv() + "." + s)
	if r == 0 {
		r = viper.GetInt("common." + s)
	}
	return r
}

func GetInt64(s string) int64 {
	//fmt.Println(GetEnv()+"."+s)
	r := viper.GetInt64(GetEnv() + "." + s)
	if r == 0 {
		r = viper.GetInt64("common." + s)
	}
	return r
}

func IsSet(s string) bool {
	if viper.IsSet(s) {
		return true
	}
	return viper.IsSet("common." + s)
}

func GetDuration(s string) time.Duration {
	return viper.GetDuration(GetConfigKey(s))
}

func GetStringSlice(s string) []string {
	return viper.GetStringSlice(GetConfigKey(s))
}

func GetIntSlice(s string) []int {
	return viper.GetIntSlice(GetConfigKey(s))
}

func GetBool(s string) bool {
	return viper.GetBool(GetConfigKey(s))
}

func GetConfigKey(s string) string {
	key := GetEnv() + "." + s
	if !viper.IsSet(key) {
		key = "common." + s
	}
	return key
}

//根据当前环境获取interface类型
func Get(s string) interface{} {
	r := viper.Get(GetEnv() + "." + s)
	if r == nil {
		r = viper.Get("common." + s)
	}
	return r
}
