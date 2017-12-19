package main

import "fmt"

type Version struct {
	Major, Minor, Patch uint //Choosing uint to avoid erroneous usage of negative value
}

func (v Version) VersionApi() {
	fmt.Printf("{\"version\": \"%d.%d.%d\"}\n", v.Major, v.Minor, v.Patch)
}

func main() {
	//Insert the hard-coded version of the code
	My_version := Version{
		Major: 2,
		Minor: 6,
		Patch: 5,
	}

	///Invoke the method VersionApi() to print the hard-coded version
	My_version.VersionApi()
}
