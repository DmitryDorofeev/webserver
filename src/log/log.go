package log

import "fmt"

func Write(message string) {
	enable := true
	if enable {
		fmt.Println(message)
	}
}
