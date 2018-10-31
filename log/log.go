/*
* Log methods 
*/
package log

import (
	"log"
	"os"
)

var (
	logfile *os.File
	err error
	errLog *log.Logger
	infoLog *log.Logger
)

func init() {
	logfile, err = os.OpenFile("./data/agenda.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatalln("Error when opening the file: agenda.log")
		os.Exit(1)
	}
	errLog = log.New(logfile, "[error]", log.Ldate | log.Ltime | log.Lshortfile)
	infoLog = log.New(logfile, "[info]", log.Ldate | log.Ltime)
}

// Write message to log, type 0 for [error], type 1 for [info]
func WriteLog(mess string, type_ int) {
	var (
		logger *log.Logger
	)
	if type_ == 0 {
		logger = errLog
	}else {
		logger = infoLog
	}

	logger.Println(mess)
}