package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil  {
    	return nil, err
    }
    return &Page{Title: title, Body: body}
}

func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
