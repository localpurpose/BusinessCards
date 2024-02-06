package utils

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func checkerr(err error) {
	if err != nil {
		log.Info(err)
	}
}
