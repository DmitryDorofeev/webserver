package headers

import (
	"net/url"
	"strings"
)

var exts = map[string]string{
	"txt":  "application/text",
	"html": "text/html",
	"json": "application/json",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"js":   "text/javascript",
	"css":  "text/css",
	"gif":  "image/gif",
	"swf":  "application/x-shockwave-flash",
}

func GetHeaderByExt(ext string) string {
	ct, ok := exts[ext]
	if ok {
		return ("Content-Type: " + ct)
	} else {
		return "Content-Type: text/html; charset=utf-8"
	}
}

func GetExtByFileName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

func ParseQueryString(query string) map[string]string {
	parts := strings.Split(query, " ")

	uri := strings.Split(parts[1], "?")

	path, _ := url.QueryUnescape(uri[0])

	return map[string]string{
		"method": parts[0],
		"path":   path,
	}
}

func IsDirectory(path string) bool {
	if path[len(path)-1:] == "/" {
		return true
	}
	return false
}
