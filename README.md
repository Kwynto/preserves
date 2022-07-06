# Preserves
These are preparations of canned code for all occasions. Might come in handy sometime.

## Detailed description

In any business, every master has a set of good tools, and coders have their own tools too.  
There are two large categories of tools: the first one we use constantly, every day and they wander from one project to another, and the second type of tools we use rarely, in cases of emergency.  
All these tools are well-established and tested, they are fast, efficient and safe.  
I decided to build my own box of tool code so that the code would be available at all times.  

You can quickly connect all the tools from this set  
> go get github.com/Kwynto/preserves

In your Go code, connect the blanks like this:  
```go
import "github.com/Kwynto/preserves"
```

When you stop using this package, do not forget to get rid of unnecessary dependencies in your project  
> go mod tidy

This package does not require separate documentation, as it is a collection of various powerful functions and each function has a description.  

If you want to increase the efficiency of your code, then you will need to take individual functions and copy them into your project.  
To do this, you will need to first copy the entire repository to yourself.  
> git clone https://github.com/Kwynto/preserves.git

This package has two files `ordinary.go` for commonly used functions and `curiosity.go` for functions that are rarely needed.  

You can also contribute to this repository through a pull request and this toolbox will be our shared one.  
If you want to add your own function to the repository, don't forget to write a test and a benchmark.  

## About the author

The author of the project is Constantine Zavezeon (Kwynto).  
You can contact the author by e-mail: kwynto@mail.ru  
The author accepts proposals for participation in open source projects,  
as well as willing to accept job offers.