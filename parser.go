package main

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(" href=\"/(?P<link>[a-zA-Z0-9-_/:.]+)\".*data-gtmc=\"search result\"[^>]*>(?P<name>[^<]*)<")

type PackageInfoParser interface {
	ParsePackageList(body string) []PackageInfo
}

type packageInfoParser struct{}

func DefaultParser() PackageInfoParser {
	p := &packageInfoParser{}
	return p
}

func (p *packageInfoParser) ParsePackageList(body string) []PackageInfo {
	matches := regex.FindAllStringSubmatch(string(body), -1)
	packages := make([]PackageInfo, 0)
	for _, match := range matches {
		link := match[1]
		name := strings.TrimSpace(match[2])
		packages = append(packages, PackageInfo{Name: name, Url: link})
	}
	return packages
}
