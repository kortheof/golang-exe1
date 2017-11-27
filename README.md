### Golang Exercises Delivery Repository

Use the following command to clone this Repository to a local directory:

```bash
$ git clone https://github.com/kortheof/golang-exercises.git
$ cd ./golang-exercises
```

# Exercise 3

In order to view the source code of the third exercise:
```bash
$ cd ./exercise3/src
```
In order to start the Web Server:
```bash
$ go run httpServer.go
```

Use the following calls, in order to check the API responses
```bash
$ curl http://localhost:8000/average
$ curl -d '{"surname":"Eksypnakias"}' http://localhost:8000/employee
```
# Exercise 1

In order to view the source code of the first exercise:

```bash
$ cd ./exercise1/src/
```

In this branch of the code, the versions are provided during compilation time.
The code has been compiled using the following command:

```bash
$ go build -o ../bin/my_version -ldflags "-X main.major=2 -X main.minor=6 -X main.patch=50" my_version.go
```

In order to execute the compiled binary, do the following:

```bash
$ cd ../bin
$ ./my_version
```

This will print the versions that were inserted during compilation.
