package status

const (
	OK           string = "200 OK"
	NOT_FOUND    string = "404 NOT FOUND"
	ERROR        string = "500 INTERNAL SERVER ERROR"
	DEFAULT_FILE string = "/index.html"
	FILE_404     string = "/404.html"
	HTTP_VERSION string = "1.1"
)


func GetStatusLine(status string) string {
	return "HTTP/" + HTTP_VERSION + " " + status + "\n"
}
