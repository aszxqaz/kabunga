package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: kabunga [go get args and flags] [query]")
		os.Exit(1)
	}

	f := DefaultFetcher()
	query := os.Args[len(os.Args)-1]
	html, err := f.FetchPackageList(query)
	if err != nil {
		fmt.Println("Error fetching package data:", err)
		os.Exit(1)
	}

	parser := DefaultParser()
	packages := parser.ParsePackageList(html)

	propter := DefaultPrompter()
	packageInfo, err := propter.Prompt(packages)
	if err != nil {
		fmt.Println("Error selecting package:", err)
		os.Exit(1)
	}

	if packageInfo != nil {
		args := os.Args[1 : len(os.Args)-1]
		parts := append([]string{"get"}, args...)
		parts = append(parts, packageInfo.Url)
		cmd := exec.Command("go", parts...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
