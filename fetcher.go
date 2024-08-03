package main

import (
	"fmt"
	"io"
	"net/http"
)

type HtmlFetcher interface {
	FetchPackageList(query string) (string, error)
}

type htmlFetcher struct {
	baseUrl string
}

func DefaultFetcher() HtmlFetcher {
	f := &htmlFetcher{baseUrl: "https://pkg.go.dev"}
	return f
}

func (f *htmlFetcher) FetchPackageList(query string) (string, error) {
	url := fmt.Sprintf("%s/search?q=%s", f.baseUrl, query)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching package data: %v\n", err)
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return "", err
	}
	return string(body), nil
}
