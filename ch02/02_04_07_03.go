package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(nil)
	}

	request.Header.Set("X-TEST", "ヘッダーも変更できます")
	request.Write(os.Stdout)
}
