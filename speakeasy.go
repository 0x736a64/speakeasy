package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func setHeadersThenServe(handler http.Handler) http.HandlerFunc {
	return func(server http.ResponseWriter, request *http.Request) {
		server.Header().Add("Cache-Control", "no-store")
		handler.ServeHTTP(server, request)
	}
}

func run(dir, port, host string) {
	open.Start(host)
	fmt.Println("-> Listening on ", host)
	fmt.Println("-> Press ctrl-c to kill process")
	http.Handle("/", setHeadersThenServe(http.FileServer(http.Dir(dir))))
	panic(http.ListenAndServe(port, nil))
}
