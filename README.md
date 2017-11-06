# golang-exercises
Golang Exercises Delivery Repository

Use the following command to clone this Repository to a local directory:

$ git clone https://github.com/kortheof/golang-exercises.git

$ cd ./golang-exercises

In order to view the source code of the first exercise:

$ cd ./exercise1/src/

In this branch of the code, the versions are provided during compilation time.
The code has been compiled using the following command:

$ go build -o ../bin/my_version -ldflags "-X main.major=2 -X main.minor=6 -X main.patch=50" my_version.go

In order to execute the compiled binary, do the following:

$ cd ../bin

$ ./my_version

This will print the versions that were inserted during compilation.
