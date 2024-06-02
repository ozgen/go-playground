package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil
	fmt.Println(string(bs))
	fmt.Println("logs bytes: ", len(bs))
	return len(bs), nil
}

func main() {

	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	//bs := make([]byte, 999999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	lw := logWriter{}
	fmt.Println("---------------------------------")
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}
