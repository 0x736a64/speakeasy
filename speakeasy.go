package main

import(
	"github.com/kardianos/osext"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"fmt"
)

func main() {
	dir, _ := osext.ExecutableFolder()
	root := http.FileServer(http.Dir(dir))
    port := ":3000"
	http.Handle("/", root)
	fmt.Println("Listening on localhost" + port)
	fmt.Println("press 'ctrl-c' to terminate process")
	open.Start("http://localhost" + port)
	http.ListenAndServe(port, nil)
}