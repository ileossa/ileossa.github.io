package quitoque

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// WriteFile data into file
func WriteFile(path string) {
	mydata := []byte("insert data")
	err := ioutil.WriteFile("myFile.data", mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

// ReadFile data into files
func ReadFile(path string) {
	data, err := ioutil.ReadFile("myFile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

// Extract data from html, with regext pattern passing
func Extract(regex, html string) {
	reg := regexp.MustCompile(regex)
	subchall := reg.FindAllStringSubmatch(html, -1)
	for _, element := range subchall {
		fmt.Println(element[1])
	}
}

// GetHTML get html
func GetHTML(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
