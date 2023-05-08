package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	str := "sample"
	fileName := str + ".pdf"
	URL := "https://www.africau.edu/images/default/sample.pdf"
	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s downlaod in current working directory", fileName)
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
