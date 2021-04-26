package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type logWritter struct{}

func main() {

	urlList := genUrl()

	c := make(chan string)

	for _, url := range urlList {
		go statusCheck(url, c)
	}

	for i := 0; i < len(urlList); i++ {
		fmt.Println(<-c)
	}

}

func statusCheck(urls string, c chan string) {

	resp, err := http.Get(urls)

	if err != nil {
		c <- "error happened with this " + urls
		fmt.Println("error triggered")
		return
	}
	// get the url and add the timestamp
	go printHtml(urls, resp.Body, c)
}

func printHtml(fileName string, response io.ReadCloser, c chan string) {
	lw := logWritter{}
	count, _ := io.Copy(lw, response)
	if count > 0 {
		c <- "File has been written"
	} else {
		c <- "File Not written"
	}
}

func (logWritter) Write(bs []byte) (int, error) {
	err := ioutil.WriteFile("HTML SOurce COde", bs, 0644)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}
