package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type logwritter struct{}

func GetHtml() io.ReadCloser {

	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		log.Fatalln("API error")
	}

	return resp.Body
}

// custom write function
func (logwritter) Write(bs []byte) (int, error) {
	//fmt.Println(string(bs))
	err := ioutil.WriteFile("google_Index_page.html", bs, 0644)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

func main() {
	// Write(p []byte) (n int, err error)
	// Read(p []byte) (n int, err error)
	// func Copy(dst Writer, src Reader) (written int64, err error)
	response := GetHtml()
	lw := logwritter{}
	io.Copy(lw, response)
}
