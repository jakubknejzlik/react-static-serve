package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func serveIndex(w http.ResponseWriter, r *http.Request, fs http.FileSystem) {
	file, err := fs.Open(path.Clean("index.html"))
	if err != nil {
		handleError(w, err)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text-html")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
}

func fileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			serveIndex(w, r, fs)
			return
		}
		if strings.HasSuffix(r.URL.Path, "index.html") {
			serveIndex(w, r, fs)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func main() {
	fs := fileServerWithCustom404(http.Dir("."))
	http.Handle("/", Gzip(fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Println("Listening..." + port)
	http.ListenAndServe(":"+port, nil)
}
