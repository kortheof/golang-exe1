package main

import (
	"fmt"
)

type Version struct {
	Major, Minor, Patch string
}

func (v Version) VersionApi() {
	fmt.Printf("{\"version\": \"%s.%s.%s\"}\n", v.Major, v.Minor, v.Patch)
}

//Variables that will hold the versions upon complilation
var (
	major, minor, patch string
)

func main() {
	//Struct of versions
	My_version := Version{major, minor, patch}

	My_version.VersionApi()
}

//go build -ldflags "-X main.major=3 -X main.minor=2 -X main.patch=1" go_exercise1.go
