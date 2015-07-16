package main

import(
	"github.com/kardianos/osext"
	"net/http"
    "strings"
	"log"
)

func main() {
	dir, _ := osext.Executable()
    dir = strings.TrimSuffix(dir, "speakeasy")
	root := http.FileServer(http.Dir(dir))
    port := ":3000"
	http.Handle("/", root)
	log.Println("listening on localhost" + port)
	http.ListenAndServe(port, nil)
}