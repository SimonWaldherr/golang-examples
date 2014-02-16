# Go Examples

## about

My first contact with [golang](http://golang.org) was in 2009, but then we went different ways. Now i want work again with golang and help others with their first contact. These examples explain the basics of golang.  
There will be more examples from time to time.

## install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```
sudo apt-get install golang
```

## examples

print Hello World with comments

```
go run HelloWorld.go
```

declare variables and print them

```
go run var.go
```

declare array and print it's items

```
go run array.go
```

do something multiple times

```
go run for.go
```

return the current time/date in various formats

```
go run time.go
```

return pseudo random integer values

```
go run random.go
```

concat strings in two different ways

```
go run cat.go
```

read via cli provided input data

```
go run args.go string string2
```

if statement in golang

```
go run if.go Hello
```

modulo operation finds the remainder of division

```
go run modulo.go
```

split a string by another string and make an array from the result

```
go run split.go
```

split a string via regular expression and make an array from the result

```
go run splitbyregex.go
```

calculate and print the fibonacci numbers

```
go run fibonacci.go
```

calculate and print the requested (32th) prime number

```
go run prime.go 32
```

pop and push in golang

```
go run lifo.go
```

convert go object to json string

```
go run json.go
```

the smallest golang http server

```
go run httpd.go
```

the smallest golang http proxy

```
go run proxy.go
```

demonstrate the power of multithreading / parallel computing  
you have to set GOMAXPROCS to something greater than 1 to see any effect

```
export GOMAXPROCS=8
time go run parallel.go true
time go run parallel.go false
```

## compile

One great aspect of golang is, that you can start go applications via ```go run name.go```, but also compile it to an executable with ```go build name.go```. After that you can start the compiled version which starts much faster.
If you start fibonacci.go and the compiled version you will notice, that the last line which contains the execution time doesn't differ much, but if you start it with ```time ./fibonacci 32``` and ```time go run ./fibonacci.go 32``` you will see the difference.

## license

The MIT License  
Copyright © 2014 Simon Waldherr  

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

