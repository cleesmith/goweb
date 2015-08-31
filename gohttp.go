// this is a more secure httpd, coz it uses:
// http.FileServer and http.Dir, which have
// file system checks to avoid break outs
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logit *log.Logger

func main() {
	port := "80"
	logit = log.New(os.Stdout,
		"",
		// log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds|log.Llongfile)
		log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds)
	fmt.Printf("Serving files in the current directory on port %v\n", port)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logit.Fatal("ListenAndServe: ", err)
	}
}
