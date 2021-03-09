package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 3. 命令行选项
// func init() {
// 	pflag.Int("redis.port", 8381, "Redis port to connect")

// 	// 绑定命令行
// 	viper.BindPFlags(pflag.CommandLine)
// }

// 4. 环境变量
// func init() {
// 	// 绑定全部环境变量
// 	// viper.AutomaticEnv()
// 	// 绑定单个环境变量
// 	viper.BindEnv("redis.port")
// 	viper.BindEnv("go.path", "GOPATH")
// }

// 6.Umarshal读取
type Config struct {
	AppName  string
	LogLevel string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

func main() {
	// // 3. 命令行选项
	// pflag.Parse
	// 1. Get读取键
	// viper.SetConfigName("config") // 设置文件名
	// viper.SetConfigType("toml")   // 配置类型
	// viper.AddConfigPath(".")      // 搜索路径
	// // viper.SetDefault("redis.port", 6379) //设置默认值
	// err := viper.ReadInConfig() // 读取配置
	// if err != nil {
	// 	log.Fatal("read config failed: %v", err)
	// }

	// // Get 读取配置值
	// fmt.Println(viper.Get("app_name"))
	// fmt.Println(viper.Get("log_level"))

	// fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	// fmt.Println("mysql port: ", viper.Get("mysql.port"))
	// fmt.Println("mysql user: ", viper.Get("mysql.user"))
	// fmt.Println("mysql password: ", viper.Get("mysql.password"))
	// fmt.Println("mysql database: ", viper.Get("mysql.database"))

	// fmt.Println("redis ip: ", viper.Get("redis.ip"))
	// fmt.Println("redis port: ", viper.Get("redis.port"))

	// // 4. 环境变量
	// fmt.Println("GOPATH: ", viper.Get("GOPATH"))
	// fmt.Println("go path: ", viper.Get("go.path"))
	// fmt.Println("redis.port: ", viper.Get("redis.port"))

	// 2. GetType读取键
	// viper.SetConfigName("config")
	// viper.SetConfigType("toml")
	// viper.AddConfigPath(".")
	// // viper.Set("redis.port", 5381)
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal("read config failed: %v", err)
	// }

	// fmt.Println("protocols: ", viper.GetStringSlice("server.protocols")) // 字符串切片
	// fmt.Println("ports: ", viper.GetIntSlice("server.ports"))            // 整型切片
	// fmt.Println("timeout: ", viper.GetDuration("server.timeout"))        // 时间类型

	// fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))  // 字符串
	// fmt.Println("mysql port: ", viper.GetInt("mysql.port")) // 整型

	// if viper.IsSet("redis.port") { // 判断是否设置键值
	// 	fmt.Println("redis.port is set")
	// } else {
	// 	fmt.Println("redis.port is not set")
	// }

	// fmt.Println("mysql settings: ", viper.GetStringMap("mysql")) // map类型返回
	// fmt.Println("redis settings: ", viper.GetStringMap("redis")) // map类型返回
	// fmt.Println("all settings: ", viper.AllSettings())           // map类型返回

	// 5. io.Reader读取配置
	// 	viper.SetConfigType("toml")
	// 	tomlConfig := []byte(`
	//   app_name = "awesome web"

	//   # possible values: DEBUG, INFO, WARNING, ERROR, FATAL
	//   log_level = "DEBUG"

	//   [mysql]
	//   ip = "127.0.0.1"
	//   port = 3306
	//   user = "dj"
	//   password = 123456
	//   database = "awesome"

	//   [redis]
	//   ip = "127.0.0.1"
	//   port = 7381
	//   `)
	// 	err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	// 	if err != nil {
	// 		log.Fatal("read config failed: %v", err)
	// 	}

	// 	fmt.Println("redis port: ", viper.Get("redis.port"))

	// 6.Unmarshal读取
	// viper.SetConfigName("config")
	// viper.SetConfigType("toml")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal("read config failed: %v", err)
	// }

	// var c Config
	// viper.Unmarshal(&c)

	// fmt.Println(c.MySQL)

	// 7.保存配置
	// viper.SetConfigName("config")
	// viper.SetConfigType("toml")
	// viper.AddConfigPath(".")

	// viper.Set("app_name", "awesome web")
	// viper.Set("log_level", "DEBUG")
	// viper.Set("mysql.ip", "127.0.0.1")
	// viper.Set("mysql.port", 3306)
	// viper.Set("mysql.user", "root")
	// viper.Set("mysql.password", "123456")
	// viper.Set("mysql.database", "awesome")

	// viper.Set("redis.ip", "127.0.0.1")
	// viper.Set("redis.port", 6381)

	// err := viper.SafeWriteConfig()
	// if err != nil {
	// 	log.Fatal("write config failed: ", err)
	// }

	// 7.监听配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))

}
