/*
* Log methods 
*/
package Log

import (
	"fmt"
	"os"
	"log"
	"time"
)

var (
	logfile *os.File
	err error
)

func InitLog() {
	logfile, err = os.OpenFile("./Log/agenda.log", os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln("Error when opening the file: agenda.log")
		os.Exit(1)
	}
	
}

// Write message to log, type 0 for [error], type 1 for [info]
func WriteLog(mess string, type_ int) {
	InitLog()
	var (
		time_ string = time.Now().Format("2006-01-02 15:04:09")
		prefix string
	)
	if type_ == 0 {
		prefix = "[Error]"
	}else {
		prefix = "[Info]"
	}
	logMess := fmt.Sprintf("%s %s %s\n", time_, prefix, mess)
	count, _ := logfile.Seek(0, os.SEEK_END)
	logfile.WriteAt([]byte(logMess), count)
	logfile.Close()
}