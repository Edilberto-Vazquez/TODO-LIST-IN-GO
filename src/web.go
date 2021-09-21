package main

import (
	"fmt"
	"io"
	"net/http"
)

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type webWriter struct {
}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("http://google.com")
	Error(err)
	e := webWriter{}
	io.Copy(e, res.Body)
}
