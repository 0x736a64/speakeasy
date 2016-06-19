package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"net/http"
  "github.com/kardianos/osext"
  "github.com/skratchdot/open-golang/open"
)

func main() {
	dir := getDir()
	port := getPort()
	host := getHost(port)
	run(dir, port, host)
}

func getDir() string {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getHost(port string) string {
	return strings.TrimSpace(fmt.Sprintf("http://localhost%s", port))
}

func getPort() string {
	portInput := flag.String("port", "3000", "an open port to run the HTTP server")
	flag.Parse()
	return fmt.Sprintf(":%s", *portInput)
}

func run(dir, port, host string) {
	open.Start(host)
	fmt.Println("-> Listening on ", host)
	fmt.Println("-> Press ctrl-c to kill process")
  http.CloseIdleConnections();
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}
