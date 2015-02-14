package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
	"strings"
)

const (
	OK string = "HTTP/1.1 200 OK\n"
	NOT_FOUND string = "HTTP/1.1 404 NOT FOUND\n"
	ERROR string = "HTTP/1.1 500 INTERNAL SERVER ERROR\n"
	DEFAULT_FILE string = "/index.html"
	FILE_404 string = "/404.html"
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
		buf = make([]byte, 1024)
		request string
	)

	_, err := conn.Read(buf)

	if err != nil {
		return
	}
	request = string(buf)
	headers := strings.Split(request, "\n")

	requestType := headers[0]
	fmt.Println(requestType)
	path := strings.Split(requestType, " ")[1]
	if (path == "/") {
		path = DEFAULT_FILE
	}

	file, err := ioutil.ReadFile("./static" + path)
	respCode = OK
	if (err != nil) {
		respCode = NOT_FOUND
		file, err = ioutil.ReadFile("./static" + FILE_404)
		checkError(err)
	} 
		
	str := string(file)
	
	var response string = respCode + "Content-Type: text/html; charset=utf-8\nServer: Veefor\n\n"

	_, err2 := conn.Write([]byte(response + str))
	checkError(err2)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
