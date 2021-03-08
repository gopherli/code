# logrus

- 日志级别：logrus.SetLevel(logrus.TraceLevel)
- 输出文件名：logrus.SetReportCaller(true)
- 添加字段：logrus.WithField和logrus.WithFields
- 重定向输出： logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
- 自定义：log := logrus.New()
- 日志格式：logrus.SetFormatter
- 第三方格式：go get github.com/antonfisher/nested-logrus-formatter
- 设置钩子：syslog
    - mgorus：将日志发送到 mongodb；
    - logrus-redis-hook：将日志发送到 redis； go get github.com/rogierlommers/logrus-redis-hook
    - logrus-amqp：将日志发送到 ActiveMQ。