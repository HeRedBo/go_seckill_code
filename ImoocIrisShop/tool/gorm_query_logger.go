package tool

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type GormQueryLogger struct {
	gormLogger *logrus.Logger
}

func NewGormQueryLogger() *GormQueryLogger {
	now := time.Now()
	logFilePath := ""

	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/runtime/logs/"
	}
	logFileName := "sql_query_" + now.Format("2006-01-02") + ".log"

	//Log file
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("logFilePath", logFilePath)
	fmt.Println("logFileName", logFileName)
	fmt.Println("fileName", fileName)
	//Write file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//Instantiation
	logger := logrus.New()

	//Set output
	logger.Out = src

	//Set log level
	logger.SetLevel(logrus.DebugLevel)

	//Format log
	//logger.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	return &GormQueryLogger{
		gormLogger: logger,
	}
}

func (logger *GormQueryLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logger.gormLogger.WithFields(
			log.Fields{
				//"module":        "gorm",
				//"type":          "sql",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"duration":      v[2],
			},
		).Info(v[3])
	case "log":
		logger.gormLogger.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
