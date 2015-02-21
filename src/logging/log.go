package logging

import (
    "fmt"
    "config"
)
func Write(message string) {
	enable := config.Get().EnableLog
	if enable {
		fmt.Println(message)
	}
}
