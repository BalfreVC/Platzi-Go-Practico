package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println(err)
	}
	e := escritorWeb{}
	io.Copy(e, res.Body)

}
