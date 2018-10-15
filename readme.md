# Go Examples

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples?ref=badge_shield)

## About

These examples explain the basics of golang. There will be more examples from time to time.

if you like, feel free to add more golang examples. Many thanks to all [contributors](https://github.com/SimonWaldherr/golang-examples/graphs/contributors).

## Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
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

```Shell
go run HelloWorld.go
```

Print Hello World with comments (shebang version)

```Shell
./HelloWorldShebang.go
```

Declare variables and print them

```Shell
go run var.go
```

Various ways (and styles) to print variables

```Shell
go run printf.go
```

If statement in golang

```Shell
go run if.go Hello
```

Declare array and print it's items

```Shell
go run array.go
```

Declare your own functions

```Shell
go run function.go
```

Do something multiple times

```Shell
go run for.go
```

Read via cli provided input data

```Shell
go run args.go string string2
```

Read via cli provided input data

```Shell
go run input.go
```

Or scan for it

```Shell
go run scan.go
```

Read named argument input data

```Shell
go run flag.go
```

Return the *working directory*

```Shell
go run dir.go
```

Return the current time/date in various formats

```Shell
go run time.go
```

Return pseudo random integer values

```Shell
go run random.go
```

Concat strings in two different ways

```Shell
go run cat.go
```

Modulo operation finds the remainder of division

```Shell
go run modulo.go
```

Split a string by another string and make an array from the result

```Shell
go run split.go
```

An example implementation of the Ackermann function

```Shell
go run ackermann.go
```

An example implementation of the Euclidean algorithm

```Shell
go run euklid.go
```

Make pipeable unix applications with os.Stdin

```Shell
go run pipe.go
```

Submit a function as argument

```Shell
go run functioncallback.go
```

A function returned by a function

```Shell
go run functionclosure.go
```

A function with an unknown amount of inputs (variadic function)

```Shell
go run functionvariadic.go
```

Empty interface as argument (You Don't Know Type)

```Shell
go run interfaces.go
```

Make structs (objects) which have functions

```Shell
go run oop.go
```

Dependency injection for easier testing

```Shell
cd beginner/di
go test
```

### Advanced

Benchmarking example (using JSON marshal and unmarshal for the sample)

From the root directory (`$GOPATH/github.com/SimonWaldherr/golang-examples`), run this command:

```Shell
go test -bench=. -benchmem advanced/json_bench/main_test.go
```

AES-GCM encryption example

```Shell
go run aesgcm.go
```

Bcrypt hashing example

Please install package golang.org/x/crypto/bcrypt before run this file by running `go get golang.org/x/crypto/bcrypt`

```Shell
go run bcrypt.go
```

Search element is exist in arrays or not

```Shell
go run in_array.go
```

Calculate triangles

```Shell
go run pythagoras.go (float|?) (float|?) (float|?)
```

Read from stdin (but don't wait for the enter key)

```Shell
go run getchar.go
```

Wait and sleep

```Shell
go run wait.go
```

Last in - first out - example

```Shell
go run lifo.go
```

Split a string via regular expression and make an array from the result

```Shell
go run regex.go
```

More advanced regex (with time and dates)

```Shell
go run regex2.go
```

Use my [golibs regex package](https://github.com/SimonWaldherr/golibs#regex-----)

```Shell
go run regex3.go
```

Calculate and print the fibonacci numbers

```Shell
go run fibonacci.go
```

Calculate and print the requested (32th) prime number

```Shell
go run prime.go 32
```

Do things with numbers, strings and switch-cases

```Shell
go run numbers.go
```

Pop and push in golang

```Shell
go run lifo.go
```

List files in working directory

```Shell
go run explorer.go
```

Start a ticker (do things periodically)

```Shell
go run ticker.go
```

Do something in case of a timeout

```Shell
go run timeout.go
```

Convert go object to json string

```Shell
go run json.go
```

Run unix/shell commands in go apps

```Shell
go run exec.go
```

Compress by pipe

```Shell
go run compress.go
```

Compress by file

```Shell
go run compress2.go
```

Run a self killing app

```Shell
go run suicide.go
```

### Expert

Calculate π with go (leibniz, euler and prime are running until you stop it via CTRL+C)

```Shell
go run pi2go.go leibniz
go run pi2go.go euler
go run pi2go.go prime
```

Convert from rgb to hsl

```Shell
go run color.go
```

Telnet with golang

```Shell
go run telnet.go
```

The smallest golang http server

```Shell
go run httpd.go
```

Secure golang http server

```Shell
go run httpsd.go
```

The smallest golang http proxy

```Shell
go run proxy.go
```

Read and write cookies

```Shell
go run cookies.go
```

Demonstrate the power of multithreading / parallel computing
you have to set GOMAXPROCS to something greater than 1 to see any effect

```Shell
export GOMAXPROCS=8
time go run parallel.go true
time go run parallel.go false
```

A dynamic amount of channels

```go
time go run dynparallel.go 8
```

Run the compiler and comment each line which contains an error

```Shell
go build gocomment.go
./gocomment go-app.go
```

Convert a image to a grayscale and to a color inverted image

```Shell
go run image.go
```

Sql (sqlite) golang example

```Shell
go run sqlite.go insert test
go run sqlite.go select
```

Public-key/asymmetric cryptography signing and validating

```Shell
go run ppk-crypto.go
```

Hashing (md5, sha) in go

```Shell
go run hashing.go
```

## Compile

One great aspect of golang is, that you can start go applications via ```go run name.go```, but also compile it to an executable with ```go build name.go```. After that you can start the compiled version which starts much faster.
If you start fibonacci.go and the compiled version you will notice, that the last line which contains the execution time doesn't differ much, but if you start it with ```time ./fibonacci 32``` and ```time go run ./fibonacci.go 32``` you will see the difference.

## License

Copyright © 2018 Simon Waldherr
Dual-licensed. See the [LICENSE](https://github.com/SimonWaldherr/golang-examples/blob/master/LICENSE) file for details.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSimonWaldherr%2Fgolang-examples?ref=badge_large)

## Support me

if you like what i do
feel free to support me

you can do so by:

* [donate via PayPal](https://www.paypal.me/SimonWaldherr "Donate to this project via PayPal.me") or [liberaPay](https://liberapay.com/SimonWaldherr/donate "Donate using Liberapay")
* buy me a beer or [Club-Mate](https://en.wikipedia.org/wiki/Club-Mate#Hacker_culture) at a conference
* give me a job where I can work on open source projects (please don't contact me via LinkedIn - please send an eMail or [contact me via twitter](http://twitter.com/SimonWaldherr) instead)
