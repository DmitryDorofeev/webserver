package statuses

const (
	OK           string = "HTTP/1.1 200 OK\n"
	NOT_FOUND    string = "HTTP/1.1 404 NOT FOUND\n"
	ERROR        string = "HTTP/1.1 500 INTERNAL SERVER ERROR\n"
	DEFAULT_FILE string = "/index.html"
	FILE_404     string = "/404.html"
)
