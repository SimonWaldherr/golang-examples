# Go Examples
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples?ref=badge_shield)


## About
 
These examples explain the basics of golang. There will be more examples from time to time.  

if you like, feel free to add more golang examples. Many thanks to all [contributors](https://github.com/SimonWaldherr/golang-examples/graphs/contributors).

## Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```
sudo apt-get install golang
```

[install golang manually](https://golang.org/doc/install)  
or  
[compile it yourself](https://golang.org/doc/install/source)  

## Examples

The examples are divided into three levels of difficulty. The [Beginner](https://github.com/SimonWaldherr/golang-examples#beginner) section contains very easy examples, starting with **Hello World** but also containing a few easy algorithms. The [Advanced](https://github.com/SimonWaldherr/golang-examples#advanced) section uses more complicated features of golang. Finally, the [Expert](https://github.com/SimonWaldherr/golang-examples#expert) section contains applications like telnet-clients or http-server (even with SSL).  
If you want even more golang examples, you can take a look at my other go repositories at github:  

* [GolangSortingVisualization](https://github.com/SimonWaldherr/GolangSortingVisualization) visualizes various sorting algorithms on the terminal or as gif
* [golang-minigames](https://github.com/SimonWaldherr/golang-minigames) currently only contains a snake clone
* [bbmandelbrot.go](https://github.com/SimonWaldherr/bbmandelbrot.go) calculates a mandelbrot fractal and saves it as png
* [golibs](https://github.com/SimonWaldherr/golibs) contains various Go packages (e.g. math, converter, stack, cli, ...)
* [cgol.go](https://github.com/SimonWaldherr/cgol.go) is conways game of life in Golang
* [micromarkdownGo](https://github.com/SimonWaldherr/micromarkdownGo) converts markdown to html (via regular expression)
* [wikiGo](https://github.com/SimonWaldherr/wikiGo) is a wiki software in Go
* [...](https://github.com/search?utf8=✓&q=user%3Asimonwaldherr&type=Repositories&ref=advsearch&l=Go)

All of them are published as free and open source software.

If all of this is even not enough for you, you can take a look at the following websites:

* [tour.golang.org](https://tour.golang.org/)
* [Go by example](https://gobyexample.com/)
* [Golang Book](http://www.golang-book.com/)
* [Go-Learn](https://github.com/skippednote/Go-Learn)

### Beginner

To execute a **golang** program, write ```go run``` at the cli followed by the name of the file.
You also can convert the file to a binary executable program by the command ```go build```.
If you know ```#!```, also known as [Shebang](https://en.wikipedia.org/wiki/Shebang_(Unix)), there is an equivalent for go: ```//usr/bin/env go run $0 $@ ; exit```

Print Hello World with comments

```
go run HelloWorld.go
```

Print Hello World with comments (shebang version)

```
./HelloWorldShebang.go
```

Declare variables and print them

```
go run var.go
```

Various ways (and styles) to print variables

```
go run printf.go
```

If statement in golang

```
go run if.go Hello
```

Declare array and print it's items

```
go run array.go
```

Declare your own functions

```
go run function.go
```

Do something multiple times

```
go run for.go
```

Read via cli provided input data

```
go run args.go string string2
```

Read via cli provided input data

```
go run input.go
```

Or scan for it

```
go run scan.go
```

Read named argument input data

```
go run flag.go
```

Return the *working directory*

```
go run dir.go
```

Return the current time/date in various formats

```
go run time.go
```

Return pseudo random integer values

```
go run random.go
```

Concat strings in two different ways

```
go run cat.go
```

Modulo operation finds the remainder of division

```
go run modulo.go
```

Split a string by another string and make an array from the result

```
go run split.go
```

An example implementation of the Ackermann function

```
go run ackermann.go
```

An example implementation of the Euclidean algorithm

```
go run euklid.go
```

Make pipeable unix applications with os.Stdin

```
go run pipe.go
```

Submit a function as argument

```
go run functioncallback.go
```

A function returned by a function

```
go run functionclosure.go
```

A function with an unknown amount of inputs (variadic function)

```
go run functionvariadic.go
```

Empty interface as argument (You Don't Know Type)

```
go run interfaces.go
```

Make structs (objects) which have functions

```
go run oop.go
```

### Advanced

AES-GCM encryption example

```
go run aesgcm.go
```

Bcrypt hashing example

Please install package golang.org/x/crypto/bcrypt before run this file by running `go get golang.org/x/crypto/bcrypt`

```
go run bcrypt.go
```

Search element is exist in arrays or not

```
go run in_array.go
```

Calculate triangles

```
go run pythagoras.go (float|?) (float|?) (float|?)
```

Read from stdin (but don't wait for the enter key)

```
go run getchar.go
```

Wait and sleep

```
go run wait.go
```

Last in - first out - example

```
go run lifo.go
```

Split a string via regular expression and make an array from the result

```
go run regex.go
```

More advanced regex (with time and dates)

```
go run regex2.go
```

Use my [golibs regex package](https://github.com/SimonWaldherr/golibs#regex-----)

```
go run regex3.go
```

Calculate and print the fibonacci numbers

```
go run fibonacci.go
```

Calculate and print the requested (32th) prime number

```
go run prime.go 32
```

Do things with numbers, strings and switch-cases

```
go run numbers.go
```

Pop and push in golang

```
go run lifo.go
```

List files in working directory

```
go run explorer.go
```

Start a ticker (do things periodically)

```
go run ticker.go
```

Do something in case of a timeout

```
go run timeout.go
```

Convert go object to json string

```
go run json.go
```

Run unix/shell commands in go apps

```
go run exec.go
```

Compress by pipe

```
go run compress.go
```

Compress by file

```
go run compress2.go
```

Run a self killing app

```
go run suicide.go
```

### Expert

Calculate π with go (leibniz, euler and prime are running until you stop it via CTRL+C)

```
go run pi2go.go leibniz
go run pi2go.go euler
go run pi2go.go prime
```

Convert from rgb to hsl

```
go run color.go
```

Telnet with golang

```
go run telnet.go
```

The smallest golang http server

```
go run httpd.go
```

Secure golang http server

```
go run httpsd.go
```

The smallest golang http proxy

```
go run proxy.go
```

Read and write cookies

```
go run cookies.go
```

Demonstrate the power of multithreading / parallel computing  
you have to set GOMAXPROCS to something greater than 1 to see any effect

```
export GOMAXPROCS=8
time go run parallel.go true
time go run parallel.go false
```

A dynamic amount of channels

```
time go run dynparallel.go 8
```

Run the compiler and comment each line which contains an error

```
go build gocomment.go
./gocomment go-app.go
```

Convert a image to a grayscale and to a color inverted image

```
go run image.go
```

Sql (sqlite) golang example

```
go run sqlite.go insert test
go run sqlite.go select
```

Public-key/asymmetric cryptography signing and validating

```
go run ppk-crypto.go
```

Hashing (md5, sha) in go

```
go run hashing.go
```

## Compile

One great aspect of golang is, that you can start go applications via ```go run name.go```, but also compile it to an executable with ```go build name.go```. After that you can start the compiled version which starts much faster.
If you start fibonacci.go and the compiled version you will notice, that the last line which contains the execution time doesn't differ much, but if you start it with ```time ./fibonacci 32``` and ```time go run ./fibonacci.go 32``` you will see the difference.

## License

Copyright © 2017 Simon Waldherr  
Dual-licensed. See the [LICENSE](https://github.com/SimonWaldherr/golang-examples/blob/master/LICENSE) file for details.  


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples?ref=badge_large)

## Support me

if you like what i do
feel free to support me

you can do so by:

* [donate via PayPal](https://www.paypal.me/SimonWaldherr "Donate to this project via PayPal.me") or [liberaPay](https://liberapay.com/SimonWaldherr/donate "Donate using Liberapay")
* buy me a beer at a conference
* give me a job where i can work on open source projects


<a target='_blank' rel='nofollow' href='https://app.codesponsor.io/link/bpWGS5JKXiHEB18kfHGr1x3q/SimonWaldherr/golang-examples'>
  <img alt='Sponsor' width='888' height='68' src='https://app.codesponsor.io/embed/bpWGS5JKXiHEB18kfHGr1x3q/SimonWaldherr/golang-examples.svg' />
</a>