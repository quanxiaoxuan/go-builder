package logs

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志配置
type Config struct {
	LogDir     string `json:"logDir" yaml:"logDir"  default:"temp/log"`                          // 日志保存文件夹
	LogName    string `json:"logName" yaml:"logName" default:"app"`                              // 日志文件名
	TimeFormat string `json:"timeFormat" yaml:"timeFormat" default:"2006-01-02 15:04:05.999999"` // 日志时间格式
	LogLevel   int    `json:"logLevel" yaml:"logLevel" default:"0"`                              // 0-debug; 1-info;2-warning; 3-error; 4-fatal; 5-panic
	SplitType  int    `json:"splitType" yaml:"splitType" default:"1"`                            // 日志切割方式  1-size,按大小切割; 2-time,按时间切割
	SplitSize  int    `json:"splitSize" yaml:"splitSize" default:"200"`                          // 切割大小(按文件大小,单位：mb)
	SplitDays  int    `json:"splitDays" yaml:"splitDays" default:"1"`                            // 切割天数(按生成时间,单位：天)
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups" default:"10"`                         // 最大备份数
}

// 初始化日志
func InitLogger(config *Config, logger *logrus.Logger) {
	_ = os.Mkdir(config.LogDir, os.ModePerm)
	logWriter := &lumberjack.Logger{
		Filename:   path.Join(config.LogDir, config.LogName+".log"),
		MaxBackups: 10,
	}

	if config.SplitType == 1 {
		logWriter.MaxSize = config.SplitSize
	} else if config.SplitType == 2 {
		logWriter.MaxAge = config.SplitDays
	} else {
		panic("config.SplitType error ")
	}

	format := &LogFormatter{config.TimeFormat}
	myHook := NewHook(WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}, format)
	logger.AddHook(myHook)
	//是否获取文件和行号
	logger.SetReportCaller(true)
	logger.SetFormatter(format)

	switch config.LogLevel {
	case 0:
		logger.SetLevel(logrus.DebugLevel)
	case 1:
		logger.SetLevel(logrus.InfoLevel)
	case 2:
		logger.SetLevel(logrus.WarnLevel)
	case 3:
		logger.SetLevel(logrus.ErrorLevel)
	case 4:
		logger.SetLevel(logrus.FatalLevel)
	case 5:
		logger.SetLevel(logrus.PanicLevel)
	}
}
