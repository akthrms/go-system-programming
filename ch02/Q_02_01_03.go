package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{"Hello": "World"}

	gzipWriter := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", " ")
	encoder.Encode(source)
	gzipWriter.Flush()
	gzipWriter.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
