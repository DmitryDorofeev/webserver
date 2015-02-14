package main

import (
	"fmt"
	"headers"
	"io/ioutil"
	"net"
	"os"
	"statuses"
	"strings"
)

func main() {

	service := ":9797"
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
		respCode string
		buf      = make([]byte, 1024)
		request  string
	)

	_, err := conn.Read(buf)

	if err != nil {
		return
	}
	request = string(buf)
	headerStrings := strings.Split(request, "\n")

	requestType := headerStrings[0]
	fmt.Println(requestType)
	path := strings.Split(requestType, " ")[1]
	if path == "/" {
		path = statuses.DEFAULT_FILE
	}
	ext := headers.GetExtByFileName(path)

	contentType := headers.GetHeaderByExt(ext)

	file, err := ioutil.ReadFile("../static" + path)
	respCode = statuses.OK
	if err != nil {
		respCode = statuses.NOT_FOUND
		file, err = ioutil.ReadFile("../static" + statuses.FILE_404)
		checkError(err)
	}

	str := string(file)

	var response string = respCode + contentType + "\nServer: Veefor\n\n"

	_, err2 := conn.Write([]byte(response + str))
	checkError(err2)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
