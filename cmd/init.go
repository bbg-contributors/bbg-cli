package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Init(themeURL string) {
	dir, _ := os.Getwd()
	stat, _ := ioutil.ReadDir(dir)
	if len(stat) != 0 {
		fmt.Println("The current directory is not empty.")
		os.Exit(1)
	}

	fmt.Println("Downloading theme from", themeURL)

	resp, err := http.Get(themeURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("A new blog has been created under", dir)
}
