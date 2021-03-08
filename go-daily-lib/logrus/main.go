package main

import (
	"io/ioutil"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

// // 8.设置钩子
// type AppHook struct {
// 	AppName string
// }

// // logrus钩子需要实现两个方法
// // type Hook interface {
// // 	Levels() []Level
// // 	Fire(*Entry) error
// //   }
// // 8.设置钩子
// func (h *AppHook) Levels() []logrus.Level {
// 	return logrus.AllLevels
// }

// // 8.设置钩子
// func (h *AppHook) Fire(entry *logrus.Entry) error {
// 	entry.Data["app"] = h.AppName
// 	return nil
// }

// 8.1将日志发送到redis
func init() {
	hookConfig := logredis.HookConfig{
		Host:     "127.0.0.1",
		Key:      "mykey",
		Format:   "v0",
		App:      "aweosome",
		Hostname: "localhost",
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func main() {
	// 1.级别
	// logrus.SetLevel(logrus.TraceLevel)

	// logrus.Trace("trace msg")
	// logrus.Debug("debug msg")
	// logrus.Info("info msg")
	// logrus.Warn("warn msg")
	// logrus.Error("error msg")
	// logrus.Fatal("fatal msg")
	// logrus.Panic("panic msg")

	// 2.输出文件名
	// logrus.SetReportCaller(true)
	// logrus.Info("info msg")

	// 3.添加字段
	// requestLogger := logrus.WithFields(logrus.Fields{
	// 	"user_id": 10010,
	// 	"ip":      "192.168.32.15",
	// })

	// requestLogger.Info("info msg")
	// requestLogger.Error("error msg")

	// 4.输出重定向
	// writer1 := &bytes.Buffer{}
	// writer2 := os.Stdout
	// writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatalf("create file log.txt failed: %v", err)
	// }

	// logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	// logrus.Info("info msg")

	// 5.自定义
	// log := logrus.New()

	// log.SetLevel(logrus.InfoLevel)
	// log.SetFormatter(&logrus.JSONFormatter{})

	// log.Info("info msg")

	// 6.日志格式
	// logrus.SetLevel(logrus.TraceLevel)
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	// logrus.Trace("trace msg")
	// logrus.Debug("debug msg")
	// logrus.Info("info msg")
	// logrus.Warn("warn msg")
	// logrus.Error("error msg")
	// logrus.Fatal("fatal msg")
	// logrus.Panic("panic msg")

	// 7.第三方格式
	// logrus.SetFormatter(&nested.Formatter{
	// 	HideKeys:    false,
	// 	FieldsOrder: []string{"component", "category"},
	// })

	// logrus.Info("info msg")

	// logrus.SetFormatter(&nested.Formatter{
	// 	HideKeys:        false,
	// 	TimestampFormat: time.RFC3339,
	// 	FieldsOrder:     []string{"name", "age"},
	// })

	// logrus.WithFields(logrus.Fields{
	// 	"name": "dj",
	// 	"age":  18,
	// }).Info("info msg")

	// 8.设置钩子
	// h := &AppHook{AppName: "awesome-web"}
	// logrus.AddHook(h)

	// logrus.Info("info msg")
	// 8.1将日志发送到redis
	logrus.Info("just some info logging...")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that",
	}).Info("additional fields are being logged as well")

	logrus.SetOutput(ioutil.Discard)
	logrus.Info("This will only be sent to Redis")

}
