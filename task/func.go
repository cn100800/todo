package task

import "fmt"

func Success(s string) {
	fmt.Printf("\033[32m %s \033[0m\n", s)
}

func Info(s string) {
	fmt.Printf("\033[0m %s \033[0m\n", s)
}

func Warning(s string) {
	fmt.Printf("\033[33m %s \033[0m\n", s)
}

func Error(s string) {
	fmt.Printf("\033[31m %s \033[0m\n", s)
}
