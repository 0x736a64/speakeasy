package main

import(
	"github.com/kardianos/osext"
	"net/http"
	"log"
)

func main() {
	dir, _ := osext.ExecutableFolder()
	root := http.FileServer(http.Dir(dir))
    port := ":3000"
	http.Handle("/", root)
	log.Println("listening on localhost" + port)
	http.ListenAndServe(port, nil)
}