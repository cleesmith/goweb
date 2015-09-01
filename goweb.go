package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logit *log.Logger

func main() {
	port := "80"
	logit = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds)
	logit.Printf("GoWeb: listening on port %v\n", port)
	// the following has file system checks to avoid break outs:
	fileHandler := http.FileServer(http.Dir("."))
	err := http.ListenAndServe(":"+port, logWrapper(fileHandler))
	if err != nil {
		logit.Fatal("ListenAndServe: ", err)
	}
}

// the following was taken from:
// https://github.com/ajays20078/go-http-logger

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	w.length = len(b)
	return w.ResponseWriter.Write(b)
}

// logWrapper logs the http status for a request and returns a
// http.HandlerFunc which is a wrapper to log the requests:
func logWrapper(handle http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		start := time.Now()
		writer := statusWriter{w, 0, 0}
		handle.ServeHTTP(&writer, request)
		end := time.Now()
		latency := end.Sub(start)
		statusCode := writer.status
		length := writer.length
		if request.URL.RawQuery != "" {
			logit.Printf("%v %s %s \"%s %s%s%s %s\" %d %d \"%s\" %v", end.Format("2006/01/02 15:04:05"), request.Host, request.RemoteAddr, request.Method, request.URL.Path, "?", request.URL.RawQuery, request.Proto, statusCode, length, request.Header.Get("User-Agent"), latency)
		} else {
			logit.Printf("%v %s %s \"%s %s %s\" %d %d \"%s\" %v", end.Format("2006/01/02 15:04:05"), request.Host, request.RemoteAddr, request.Method, request.URL.Path, request.Proto, statusCode, length, request.Header.Get("User-Agent"), latency)
		}
	}
}
