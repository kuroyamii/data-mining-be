package logger

import (
	"datamining-be/pkg/utilities"
	"fmt"
	"time"
)

type LogWritter struct {
}

func (lw LogWritter) Write(bytes []byte) (int, error) {
	return fmt.Printf("[%v] [%v] %v", utilities.Info(time.Now().Format("2006/01/02")), utilities.Info(time.Now().Format("15:04:05")), string(bytes))
}
