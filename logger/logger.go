package logger

import (
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/kyani-inc/logger"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
}

func Basic() Logger {
	a := logrus.New()
	a.Out = os.Stdout

	return a
}

func File(file string) Logger {
	// `f` will be managed by logrus
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return Basic()
	}

	a := logrus.New()
	a.Out = f

	return a
}

func Papertrail(host, port, appname string) Logger {
	p, _ := strconv.Atoi(port)

	return logger.New(logger.Config{
		Appname: appname,
		Host:    host,
		Port:    p,
	})
}
