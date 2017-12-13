package main

import (
	"os"
	"net/http"
	"bytes"
	"fmt"
	"log"
)

func newfileUploadRequest(path string, uri string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/octet-stream")
	return req, err
}

func main() {
	args := os.Args[1:]
	for i, _ := range args {
		if ((i==0) && (len(args)==2)) {
			request, err := newfileUploadRequest(args[i], args[i+1])
			if err != nil {
				log.Fatal(err)
			}
			client := &http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				log.Fatal(err)
			} else {
				body := &bytes.Buffer{}
				_, err := body.ReadFrom(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				resp.Body.Close()
				fmt.Println(resp.StatusCode)
				fmt.Println(resp.Header)
			}
		}

	}
}

