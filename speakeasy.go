package main

import(
	"log"
	"os"
	"net/http"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	port := ":3000"
	index := http.FileServer(http.Dir(dir))
	http.Handle("/", index)
	log.Println("Serving index from ", dir)
	http.ListenAndServe(port, nil)
}