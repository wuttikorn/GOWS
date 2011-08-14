package main

import (
	
	"http"
	"io/ioutil"
	"os"
	"fmt"
	"flag"
	"strconv"
)

var port *int = flag.Int("port")

type Page struct {
	Body  []byte
}

func loadPage(w http.ResponseWriter, r *http.Request)(*Page, os.Error) {
	title := r.URL.Path[1:]
	if title == "" {
		title = "index.html"
	}
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body}, nil
}
