package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
	"net/http"
	"github.com/kardianos/osext"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	port := setPort()
	url := setURL(port)
	dir, _ := osext.ExecutableFolder()
	root := http.FileServer(http.Dir(dir))
	http.Handle("/", root)
	serve(url, port)
}

func serve(url, port string) {
	fmt.Println("-> Listening on ", url)
	fmt.Println("-> Press ctrl-c to kill process")
	open.RunWith(url, "Google chrome")
	http.ListenAndServe(port, nil)
}

func setPort() string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a port: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(fmt.Sprintf(":%s", input))
}

func setURL(port string) string{
	return strings.TrimSpace(fmt.Sprintf("http://localhost%s", port))
}