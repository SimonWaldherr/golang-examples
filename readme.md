# Go Examples

[![Gittip donate button](http://img.shields.io/gittip/bevry.png)](https://www.gittip.com/SimonWaldherr/ "Donate weekly to this project using Gittip") [![Flattr donate button](https://raw.github.com/balupton/flattr-buttons/master/badge-89x18.gif)](https://flattr.com/submit/auto?user_id=SimonWaldherr&url=http%3A%2F%2Fgithub.com%2FSimonWaldherr%2Fgolang-examples "Donate monthly to this project using Flattr")


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

###beginner

print Hello World with comments

```
go run HelloWorld.go
```

declare variables and print them

```
go run var.go
```

if statement in golang

```
go run if.go Hello
```

declare array and print it's items

```
go run array.go
```

declare your own functions

```
go run function.go
```

do something multiple times

```
go run for.go
```

read via cli provided input data

```
go run args.go string string2
```

read via cli provided input data

```
go run input.go
```

return the *working directory*

```
go run dir.go
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

modulo operation finds the remainder of division

```
go run modulo.go
```

split a string by another string and make an array from the result

```
go run split.go
```

###advanced

calculate triangles

```
go run pythagoras.go (float|?) (float|?) (float|?)
```

wait and sleep

```
go run wait.go
```

last in - first out - example

```
go run lifo.go
```

split a string via regular expression and make an array from the result

```
go run regex.go
```

calculate and print the fibonacci numbers

```
go run fibonacci.go
```

calculate and print the requested (32th) prime number

```
go run prime.go 32
```

do things with numbers, strings and switch-cases

```
go run numbers.go
```

pop and push in golang

```
go run lifo.go
```

convert go object to json string

```
go run json.go
```

run unix/shell commands in go apps

```
go run exec.go
```

run a self killing app

```
go run suicide.go
```

###expert

convert from rgb to hsl

```
go run color.go
```

telnet with golang

```
go run telnet.go
```

the smallest golang http server

```
go run httpd.go
```

secure golang http server

```
go run httpsd.go
```

the smallest golang http proxy

```
go run proxy.go
```

read and write cookies

```
go run cookies.go
```

demonstrate the power of multithreading / parallel computing  
you have to set GOMAXPROCS to something greater than 1 to see any effect

```
export GOMAXPROCS=8
time go run parallel.go true
time go run parallel.go false
```

a dynamic amount of channels

```
time go run dynparallel.go 8
```

run the compiler and comment each line which contains an error

```
go build gocomment.go
./gocomment go-app.go
```

convert a image to a grayscale and to a color inverted image

```
go run image.go
```

sql (sqlite) golang example

```
go run sqlite.go insert test
go run sqlite.go select
```

public-key/asymmetric cryptography signing and validating

```
go run ppk-crypto.go
```

hashing (md5, sha) in go

```
go run hashing.go
```

## compile

One great aspect of golang is, that you can start go applications via ```go run name.go```, but also compile it to an executable with ```go build name.go```. After that you can start the compiled version which starts much faster.
If you start fibonacci.go and the compiled version you will notice, that the last line which contains the execution time doesn't differ much, but if you start it with ```time ./fibonacci 32``` and ```time go run ./fibonacci.go 32``` you will see the difference.

## license

Copyright © 2014 Simon Waldherr  
dual-licensed  

###MIT

The [MIT License](http://opensource.org/licenses/MIT)  

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

###CC

[cc by sa](http://creativecommons.org/licenses/by-sa/4.0/)  

You are free to:  

**Share** — copy and redistribute the material in any medium or format  
**Adapt** — remix, transform, and build upon the material  
for any purpose, even commercially.  
The licensor cannot revoke these freedoms as long as you follow the license terms.  

Under the following terms:  

**Attribution** — You must give **appropriate credit**, provide a link to the license, and **indicate if changes were made**. You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.  
**ShareAlike** — If you remix, transform, or build upon the material, you must distribute your contributions under the **same license** as the original.  

**No additional restrictions** — You may not apply legal terms or **technological measures** that legally restrict others from doing anything the license permits.

