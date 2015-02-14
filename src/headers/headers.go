package headers

import (
	"strings"
)

var exts = map[string]string{
	"txt":  "application/text",
	"html": "text/html",
	"json": "application/json",
	"jpg":  "image/jpeg",
}

func GetHeaderByExt(ext string) string {
	ct, ok := exts[ext]
	if ok {
		return ("Content-Type: " + ct + "; charset=utf-8")
	} else {
		return "Content-Type: text/html; charset=utf-8"
	}
}

func GetExtByFileName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}
