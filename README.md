GO tutorial
===========

# Setup

https://go.dev/doc/install

# Environment


GOPATH -
# Project Structure

```
src/        # source files

bin/        # executable commands
```



# Commands

```go
go version

go run first.go
```

# Packages

https://pkg.go.dev/std

# Concurrency
Go supports concurrent execution of tasks. It means Go can execute multiple tasks simultaneously. It is different from the concept of parallelism. In parallelism, a task is split into small subtasks and are executed in parallel. But in concurrency, multiple tasks are being executed simultaneously. Concurrency is achieved in Go using Goroutines and Channels.

Goroutines
A goroutine is a function which can run concurrently with other functions. Usually when a function is invoked the control gets transferred into the called function, and once its completed execution control returns to the calling function. The calling function then continues its execution. The calling function waits for the invoked function to complete the execution before it proceeds with the rest of the statements.

But in the case of goroutine, the calling function will not wait for the execution of the invoked function to complete. It will continue to execute with the next statements. You can have multiple goroutines in a program.

Also, the main program will exit once it completes executing its statements and it will not wait for completion of the goroutines invoked.

## Channels
Channels are a way for functions to communicate with each other. It can be thought as a medium to where one routine places data and is accessed by another routine in Golang server.

A channel can be declared with the syntax

channel_variable := make(chan datatype)
Example:

	ch := make(chan int)
	
### Sending
```
channel_variable <- variable_name	
```

### Receiving


 variable_name := <- channel_variable


# Packaging Local


WORKS

see https://go.dev/doc/tutorial/create-module

https://go.dev/doc/tutorial/call-module-code

Use :

```go
go mod ...
go mod tidy

```
Example:
```
/home/user/sdk/go1.18.4

export GOPATH=~/git/MyGo
```

To install custom package, see "calculation" directory

1) create a package
1) ```cd calculation```
1) ```sudo go install```


## Module


Using modules
```
go mod init calculation
```

go mod tidy
This command helps to remove the unwanted dependencies. Go build would only update the dependencies and don't touch the ones which are no more needed. Tidy will help you clean ;)
go mod vendor
If you want to have a vendor folder, then just run this. This command downloads all dependencies defined in go.mod and places it inside the /vendor in your root. There are cases when the robustness of the system is a priority and we want our dependencies to present beforehand.

##  Multi-module

https://www.sobyte.net/post/2022-01/go-multi-module/


* Executable commands must always use package main.


# Defer
Defer statements are used to defer the execution of a function call until the function that contains the defer statement completes execution.
Stacking defer is using multiple defer statements. Suppose you have multiple defer statements inside a function. Go places all the deferred function calls in a stack, and once the enclosing function returns, the stacked functions are executed in the Last In First Out(LIFO) order. 

# References

https://go.dev/doc/gopath_code
https://www.guru99.com/google-go-tutorial.html
https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html#gopath

https://docs.cloudfoundry.org/buildpacks/go/index.html

https://github.com/SimonWaldherr/golang-examples

https://insujang.github.io/2020-04-04/go-modules/

https://go.dev/doc/code

## Useful modules

LOG - https://github.com/rs/zerolog
Web Framework - https://github.com/gin-gonic/gin
Swagger - https://github.com/swaggo/swag
 - https://github.com/getkin/kin-openapi
 
 
