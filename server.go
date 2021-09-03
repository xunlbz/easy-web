package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port    int
	help    bool
	rootDir string
)

func init() {

	flag.IntVar(&port, "port", 8080, "set web port,default is 8080")
	flag.BoolVar(&help, "help", false, "help")
	flag.StringVar(&rootDir, "root_dir", "./html", "html dir")

}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	http.Handle("/", http.FileServer(http.Dir(rootDir)))
	fmt.Printf("web server listen to http://localhost:%d \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
