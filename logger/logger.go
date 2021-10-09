package logger

import (
	"fmt"
	"os"
	"time"
)

func LogError(error interface{}) {
	Log(fmt.Sprintf("[%v ERROR] %v\n", time.Now().Format("02-01-2006 15:04:05"), error))
}

func LogWarning(warning interface{}) {
	Log(fmt.Sprintf("[%v WARNING] %v\n", time.Now().Format("02-01-2006 15:04:05"), warning))
}

func LogInfo(info interface{}) {
	Log(fmt.Sprintf("[%v INFO] %v\n", time.Now().Format("02-01-2006 15:04:05"), info))
}

func Log(msg string) {
	file, err := os.OpenFile("log/"+time.Now().Format("02_01_2006")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[%v ERROR] Error (%v) occured while logging %v\n", time.Now().Format("02-01-2006 15:04:05"), err, msg)
		}

		if err = file.Close(); err != nil {
			fmt.Printf("[%v ERROR] Error (%v) occured while closing log file\n", time.Now().Format("02-01-2006 15:04:05"), err)
		}
	}()

	if err != nil {
		panic(err)
	}

	fmt.Printf(msg)

	_, err = fmt.Fprintf(file, msg)
	if err != nil {
		panic(err)
	}
}
