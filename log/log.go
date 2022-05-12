package log

// by gee :https://geektutu.com/post/geeorm-day1.html

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	// Lshortfile show filename & codeline
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(os.Stdout, "\033[33m[warn ]\033[0m ", log.LstdFlags|log.Lshortfile)
	successLog = log.New(os.Stdout, "\033[36m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog,warnLog,successLog}
	mutex sync.Mutex
)

var (
	// ?Fatal
	Error = errorLog.Println
	Info = infoLog.Println
	Warn = warnLog.Println
	Success = successLog.Println
)

const (
	InfoLevel = iota	//0
	ErrorLevel			//1
	Disabled			//2
)

func SetLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}