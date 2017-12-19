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
