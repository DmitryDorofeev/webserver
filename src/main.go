package main

import (
	"config"
	"fmt"
	"headers"
	"io/ioutil"
	"logging"
	"net"
	"os"
	"status"
	"strings"
)

const (
	STRING_SEPARATOR string = "\n"
)

func main() {

	service := fmt.Sprintf(":%v", config.Get().Port)
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var (
		respCode      string
		buf           = make([]byte, 1024)
		HttpRequest   string
		contentType   string
		str           string
		contentLength string
		isDirectory   bool = false
		isHeadRequest bool = false
	)

	_, err := conn.Read(buf)

	if err != nil {
		return
	}

	HttpRequest = string(buf)
	headerStrings := strings.Split(HttpRequest, STRING_SEPARATOR)

	requestString := headerStrings[0] // request string - first string

	request := headers.ParseQueryString(requestString) // TODO: request must be struct

	path := request["path"]

	if request["method"] == "POST" {

		str = ""
		respCode = status.GetStatusLine(status.BAD_REQUEST)
		contentType = ""
		contentLength = ""

	} else if request["method"] == "HEAD" {

		isHeadRequest = true

		if headers.IsDirectory(path) {
			path += config.Get().DefaultFile
		}

		ext := headers.GetExtByFileName(path)

		contentType = headers.GetHeaderByExt(ext)

		file, err := ioutil.ReadFile(config.Get().Root + path)

		logging.Write(config.Get().Root + path)

		respCode = status.GetStatusLine(status.OK)

		if err != nil {
			respCode = status.GetStatusLine(status.NOT_FOUND)
			file, err = ioutil.ReadFile(config.Get().Root + status.FILE_404)
			checkError(err)
		}

		str = string(file)

		contentLength = fmt.Sprintf("Content-Length: %v\r\n", len(str))

		str = "\r\n\r\n"

	} else {

		if headers.IsDirectory(path) {
			path += config.Get().DefaultFile
			isDirectory = true
		}

		ext := headers.GetExtByFileName(path)

		contentType = headers.GetHeaderByExt(ext)

		file, err := ioutil.ReadFile(config.Get().Root + path)

		logging.Write(config.Get().Root + path)

		respCode = status.GetStatusLine(status.OK)
		if err != nil {
			if isDirectory {
				respCode = status.GetStatusLine(status.FORBIDDEN)
				file = []byte("Forbidden")
			} else {
				respCode = status.GetStatusLine(status.NOT_FOUND)
				file, err = ioutil.ReadFile(config.Get().Root + status.FILE_404)
			}

		}

		if strings.Contains(path, "..") {
			respCode = status.GetStatusLine(status.FORBIDDEN)
			file = []byte("Forbidden")
		}

		str = string(file)

		contentLength = fmt.Sprintf("Content-Length: %v\r\n", len(str))

	}

	var response string = respCode + contentType + "\r\nConnection: close\r\nServer: DmitryDorofeevAwesomeServer\r\n"
	response += contentLength
	response += "\r\n"
	if !isHeadRequest {
		response += str
	}

	_, err2 := conn.Write([]byte(response))
	checkError(err2)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
