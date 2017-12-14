package main

import (
	"net/http"
	"io"
	"fmt"
	"log"
	"os"

)

func receiveHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Starting receiver....")
	file, header, err := r.FormFile("file")


	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err=", err)
		os.Exit(1)
	}
	out, err := os.Create(dir + "/" + header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully: ")
	fmt.Fprintf(w, header.Filename)
}



func main() {

	http.HandleFunc("/receive", receiveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))


}
