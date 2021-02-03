package logging

import (
	"fmt"
	"log"
	"os"

	"github.com/notobject/sessiond/pkg/config"
)

// NewLogger 创建日志对象
func NewLogger(prefix string, c config.LogConfig) (*log.Logger, error) {
	outFile, err := os.OpenFile(c.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed! err=", err)
		return nil, err
	}
	logger := log.New(outFile, "[session] ", log.Lshortfile|log.Ldate|log.Ltime)
	return logger, nil
}
