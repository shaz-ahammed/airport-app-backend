package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"time"
)

const Banner = `
 ____   ____ ____  ____    ___ ____  
|    \ /    |    \|    \  /  _]    \ 
|  o  )  o  |  _  |  _  |/  [_|  D  )
|     |     |  |  |  |  |    _]    / 
|  O  |  _  |  |  |  |  |   [_|    \ 
|     |  |  |  |  |  |  |     |  .  \
|_____|__|__|__|__|__|__|_____|__|\_|
    
`

const LogFilePath = "logs/misc.log"

func main() {
	// Display banner
	fmt.Print(Banner)

	// Setup logger
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   LogFilePath,
		MaxSize:    5,  // Max megabytes before log is rotated
		MaxBackups: 90, // Max number of old log files to keep
		MaxAge:     60, // Max number of days to retain log files
		Compress:   true,
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC1123Z})

	logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
	log.SetOutput(logMultiWriter)

	log.WithFields(log.Fields{
		"Runtime version": runtime.Version(),
		"Number of CPUs":  runtime.NumCPU(),
		"Arch":            runtime.GOARCH,
	}).Info("Application initializing")
}
