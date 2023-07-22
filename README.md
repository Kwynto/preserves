# Preserves
These are preparations of canned code for all occasions. Might come in handy sometime.  
[![GoDoc](https://godoc.org/github.com/Kwynto/preserves?status.svg)](https://godoc.org/github.com/Kwynto/preserves)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kwynto/preserves)](https://goreportcard.com/report/github.com/Kwynto/preserves)
[![GitHub](https://img.shields.io/github/license/Kwynto/preserves)](https://github.com/Kwynto/preserves/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/Kwynto/preserves/branch/main/graph/badge.svg?token=077VB7XVQN)](https://codecov.io/gh/Kwynto/preserves) 

**_This repository is under development._**

## Detailed description

In any business, every master has a set of good tools, and coders have their own tools too.  
There are two large categories of tools: the first one we use constantly, every day and they wander from one project to another, and the second type of tools we use rarely, in cases of emergency.  
All these tools are well-established and tested, they are fast, efficient and safe.  
I decided to build my own box of tool code so that the code would be available at all times.  

You can quickly connect all the tools from this set  
> go get github.com/Kwynto/preserves

In your Go code, connect the blanks like this:  
```go
import "github.com/Kwynto/preserves/ordinary"
``` 
or 
```go
import "github.com/Kwynto/preserves/curiosity"
``` 
and more.

When you stop using this package, do not forget to get rid of unnecessary dependencies in your project  
> go mod tidy

This package does not require separate documentation, as it is a collection of various powerful functions and each function has a description.  

If you want to increase the efficiency of your code, then you will need to take individual functions and copy them into your project.  
To do this, you will need to first copy the entire repository to yourself.  
> git clone https://github.com/Kwynto/preserves.git

There are two main files in this repository `./ordinary.go` for frequently used functions and `./curiosity.go` for rarely used functions, as well as other packages.

You can also contribute to this repository through a pull request and this toolbox will be our shared one.  
If you want to add your own function to the repository, don't forget to write a test and a benchmark.  
You can create your own package like `./pkg/<your_name>`.  

**Check out the documentation**

Look at the documentation in two steps.  
First, in the console, run:
> godoc -http=:8080

And then in your web browser navigate to the uri:
> http://localhost:8080

*The `godoc` utility may not be present in your Go build and you may need to install it  
command `go get -v golang.org/x/tools/cmd/godoc`*

You can also use Go's standard functionality to view documentation in the console via `go doc`.  
For example:  
> go doc GenerateId

If your IDE is good enough, then the documentation for functions and methods will be available from your code editor.

**Testing**

Run tests:
> go test ordinary.go ordinary_test.go -v
or 
> go test curiosity.go curiosity_test.go -v

Run tests showing code coverage:
> > go test ordinary.go ordinary_test.go -v -cover
or 
> go test curiosity.go curiosity_test.go -v -cover

You can view code coverage in detail in your web browser.  
To do this, you need to sequentially execute two commands in the console:
> go test -coverprofile="coverage.out" -v  
> go tool cover -html="coverage.out"

**Performance**

You can look at code performance tests:
> go test -benchmem -bench="." curiosity.go curiosity_test.go 
or 
> go test -benchmem -bench="." ordinary.go ordinary_test.go

**[⬆ back to top](#preserves)**

## About the author

The author of the project is Constantine Zavezeon (Kwynto).  
You can contact the author by e-mail: kwynto@mail.ru  
The author accepts proposals for participation in open source projects,  
as well as willing to accept job offers.
If you want to offer me a job, then first I ask you to read [this](https://github.com/Kwynto/Kwynto/blob/main/offer.md).

**[⬆ back to top](#preserves)**
