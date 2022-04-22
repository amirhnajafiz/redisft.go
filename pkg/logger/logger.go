package logger

import (
	"fmt"
	"log"
)

func Error(e error) {
	_ = fmt.Errorf("[ERROR] %s\n", e.Error())
}

func Fatal(e error) {
	log.Fatalf("[FATAL] %s\n", e.Error())
}

func Info(s string, m string) {
	fmt.Printf("[%s] %s\n", s, m)
}
