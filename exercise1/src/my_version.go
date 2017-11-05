package main

import (
	"flag"
	"fmt"
)

type Version struct {
	Major, Minor, Patch uint //Choosing uint to avoid erroneous usage of negative value
}

func (v Version) VersionApi() {
	fmt.Printf("{\"version\": \"%d.%d.%d\"}\n", v.Major, v.Minor, v.Patch)
}

func main() {

	//Declare a pointer to the Version struct to receive the command-line inserted versions
	valuePtr := &Version{}

	flag.UintVar(&valuePtr.Major, "major", 0, "Major Version of the Code")
	flag.UintVar(&valuePtr.Minor, "minor", 0, "Minor Version of the Code")
	flag.UintVar(&valuePtr.Patch, "patch", 0, "Patch Version of the Code")

	flag.Parse()

	//Invoke the method VersionAPi() to print the command-line inserted versions
	valuePtr.VersionApi() // the same as (*valuePtr).VersionApi()
}
