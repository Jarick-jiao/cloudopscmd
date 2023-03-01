/**
* @description :
* @author : Jarick
* @Date : 2021-08-27
 */

package util

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRecodeGinLog() {
	logDir := "data/log/"
	logfileName := logDir + "log_" + time.Now().Format("2006-01-02-15-04")

	// G304: File path provided as taint input
	// logfileName = filepath.Clean(logfileName)
	// if !strings.HasPrefix(logfileName, "log_") {
	// 	panic(fmt.Errorf("Unsafe input"))
	// }

	// log pathlogfileName is err
	_, errStat := os.Stat(logDir)
	if errStat != nil {
		log.Println("log dir failed, err:", errStat)
		errMkdir := os.MkdirAll(logDir, os.ModeDir)
		if errMkdir != nil {
			log.Println("mkdir log dir failed, err:", errMkdir)
		}
	}

	// G304 (CWE-22): Potential file inclusion via variable (Confidence: HIGH, Severity: MEDIUM)
	// logFile, err := os.OpenFile(logfileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	logFile, err := os.OpenFile(filepath.Clean("./"+logfileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Println("open log file failed, err:", err)
		return
	}
	// defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile) //tty and logs
	log.SetOutput(mw)
	log.SetPrefix("")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout) //output gin logs
	// 强制日志颜色化
	gin.ForceConsoleColor() // 在主程序声明
}
