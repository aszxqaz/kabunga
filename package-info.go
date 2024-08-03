package main

type PackageInfo struct {
	Name string
	Url  string
}

func (i PackageInfo) Title() string       { return i.Name }
func (i PackageInfo) Description() string { return i.Url }
func (i PackageInfo) FilterValue() string { return i.Name }
