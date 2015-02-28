package main

import (
	"fmt"
	"headers"
	"io/ioutil"
	"net"
	"os"
	"status"
	"strings"
	"logging"
	"config"
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
		respCode string
		buf      = make([]byte, 1024)
		HttpRequest  string
		contentType  string
        str string
	)

	_, err := conn.Read(buf)

	if err != nil {
		return
	}
	HttpRequest = string(buf)
	headerStrings := strings.Split(HttpRequest, STRING_SEPARATOR)

	requestString := headerStrings[0]
    
    request := headers.ParseQueryString(requestString)
    
	path := strings.Split(requestString, " ")[1]
    
    if (request["method"] == "POST") {
        
        str = ""
        respCode = status.GetStatusLine(status.BAD_REQUEST)
        contentType = ""
        
    } else {
        
        if path == "/" {
            path = "/" + config.Get().DefaultFile
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
        
    }
    
	var response string = respCode + contentType + "\nConnection: close\nServer: DmitryDorofeevAwesomeServer\n"
    response += fmt.Sprintf("Content-Length: %v\n", len(str))

    response += "\n"
	_, err2 := conn.Write([]byte(response + str))
	checkError(err2)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
