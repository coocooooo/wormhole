package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start server at http://localhost:9999")
	f, _ := os.Open("./index.html")
	defer f.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = f.Seek(0, 0)
		io.Copy(w, f)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		multipartReader, err := r.MultipartReader()
		if err != nil {
			fmt.Println(err)
		}
		p, err := multipartReader.NextPart()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("upload", p.FileName())
		temp, _ := os.Create("./" + p.FileName())
		io.Copy(temp, p)
		w.Write([]byte("all down"))
	})
	http.ListenAndServe(":9999", nil)
}
