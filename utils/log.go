package utils

import (
	//"../beego/logs"
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func init(){
	Log_init("","")
}

func Log_init(v string, l string) {
	log = logs.NewLogger(10000)
	if l == "" {
		l = "7"
	}
	log.SetLogger(logs.AdapterConsole, `{"level":`+l+`}`) //2Critical 3Error 4Warn 5 6Info 7Debug
	if v != "" {
		st := `{"filename":"` + v + `","level":4}`
		log.SetLogger(logs.AdapterFile, st)
	}
	//log.Async()
	log.Info("Log Model Init sucess!")
}

func LogDebug(format string, v ...interface{}) {
	log.Debug(format, v...)
}

func LogErr(format string, v ...interface{}) {
	log.Error(format, v...)
}
func LogInfo(format string, v ...interface{}) {
	log.Info(format, v...)
}
func LogWarn(format string, v ...interface{}) {
	log.Warn(format, v...)
}
func LogCheckErr(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

// Log function

//func LogErr(v ...interface{}) {
//
//	logfile := os.Stdout
//	log.Println(v...)
//	logger := log.New(logfile,"\r\n",log.Llongfile|log.Ldate|log.Ltime);
//	logger.SetPrefix("[Error]")
//	logger.Println(v...)
//	defer logfile.Close();
//}
//
//func LogInfo(v ...interface{}) {
//
//	logfile := os.Stdout
//	log.Println(v...)
//	logger := log.New(logfile,"\r\n",log.Ldate|log.Ltime);
//	logger.SetPrefix("[Info]")
//	logger.Println(v...)
//	defer logfile.Close();
//}
//
//func LogDebug(v ...interface{}) {
//	logfile := os.Stdout
//	log.Println(v...)
//	logger := log.New(logfile,"\r\n",log.Ldate|log.Ltime);
//	logger.SetPrefix("[Debug]")
//	logger.Println(v...)
//	defer logfile.Close();
//}
//
//func CheckError(err error) {
//	if err != nil {
//		LogErr(os.Stderr, "Fatal error: %s", err.Error())
//	}
//}
