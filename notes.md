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


#### Day01

* variable declarations: __var name dataType = value__ or __name := value__
* int, bool, float, string
* Local variables should be used
* Global variables need not be used; Shortcut declaration not allowed in global variables
* Conditionals: if-else and switch-case
* if-else paranthesis not required and positioning of curly braces important
* You can declare variables and use them in if-else

``` go
if x := 3; x % 2 == 0; {

} else {

} 
```

* switch-case does not require __break__ statements like you have in other languages
* __fallthrough__ statement to continue the execution of the next case
* case can have multiples values to match __case "a", "b", "c" :__
* __const__ keyword used to specify constants
* const variables can be grouped
* Only one keyword for loops: __for__ Used as a while loop as well
* __break__ can be used to break out of the for loop
* Packages like __strings__, __strconv__ for string operations
* strings.Replace, strings.Split, strings.ToUpper
* len("string"), strconv.AtoI, strconv.ItoA
* Loop a string from 0 to len - 1 and access the characters

#### Day02

##### Arrays 

* __[size]dataType__
* __var numbers [10]int__
* __numbers := []int {1, 2, 3}
* len(arr), copy, append
* Slices __numbers[0:2], numbers[3:], numbers[:5]__
* __make([]int, size)__ used to dynamically construct arrays

##### Maps: Key-value pairs

* __map[keyDataType]valueDataType__
* __var romanNumbers map[int]String = make(map[int]String)__
* __delete(romanNumbers, key)__
* __range__ keyword in for loop to iterate collections


























