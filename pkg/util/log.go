package util

import (
	"os"

	"github.com/withmandala/go-log"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stderr).WithColor()
}
