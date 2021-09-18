package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	// 查找配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("conf")
	// 读取配置文件信息
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	// 将数据解析到 Conf
	if err = viper.Unmarshal(Conf); err != nil {
		return
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 如果有变化，执行期内函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		if err = viper.Unmarshal(Conf); err != nil {
			return
		}
	})
	//fmt.Printf("%+v\n", Conf)
	//fmt.Printf("Log：%+v\n", Conf.LogConfig)
	//fmt.Printf("mysql：%+v\n", Conf.MySQLConfig)
	//fmt.Printf("redis：%+v\n", Conf.RedisConfig)
	return nil
}
