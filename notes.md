#### Basic thoughts on Go

* 25 keywords
* C++
* Compiled; (no JVM like entity); compiled to machine language
* Docker, K8, Terraform
* Microservices written in GO
* statically typed language, Garbage collected
* Not an OO language
* Procedural and slightly functional programming language
* Pointers in Go. No Pointer arithmetic. Not like Pointers in C, but similar to references in Java
* Data and behavior are scattered throughout the program

* https://golang.org
* __go version__
* 1.16 is the latest

#### Hello world

* hello.go

``` go
package main

import "fmt"

func main() {
  fmt.Println("Hello Go!")
}

```

* Execution starts with __main()__
* __go run hello.go__
* Create executables by running __go build hello.go__
* The executables contain the complete Go runtime, so large in size

* the file names are __snake_case___
* Go application is made up of packages. Each package contains many functions
* __main function__ in the __main package__ is the starting point












