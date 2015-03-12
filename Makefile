all: 
	rm -f ./httpd
	export GOPATH=${PWD}
	go build -o ./httpd ./src/main.go