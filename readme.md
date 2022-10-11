# Go Examples

[![DOI](https://zenodo.org/badge/9459712.svg)](https://zenodo.org/badge/latestdoi/9459712)  
[![Go Report Card](https://goreportcard.com/badge/github.com/simonwaldherr/golang-examples)](https://goreportcard.com/report/github.com/simonwaldherr/golang-examples)  
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)  

NEW: now with an [online live editor with Golang support](https://simonwaldherr.github.io/golang-examples/). Try out and edit the examples directly in the browser:
[![SimonWaldherr/golang-benchmarks Online Editor](https://simonwaldherr.github.io/golang-examples/golang-examples.png)](https://simonwaldherr.github.io/golang-examples/)  



If you liked this project, you may also like my [golang-benchmarks](https://github.com/SimonWaldherr/golang-benchmarks) repository:  
[![SimonWaldherr/golang-benchmarks - GitHub](https://gh-card.dev/repos/SimonWaldherr/golang-benchmarks.svg?fullname)](https://github.com/SimonWaldherr/golang-benchmarks)  
my [sql-examples](https://github.com/SimonWaldherr/sql-examples) repository:  
[![SimonWaldherr/sql-examples - GitHub](https://gh-card.dev/repos/SimonWaldherr/sql-examples.svg?fullname)](https://github.com/SimonWaldherr/sql-examples)  
or my [rpi-examples](https://github.com/SimonWaldherr/rpi-examples) repository:  
[![SimonWaldherr/rpi-examples - GitHub](https://gh-card.dev/repos/SimonWaldherr/rpi-examples.svg?fullname)](https://github.com/SimonWaldherr/rpi-examples)

## About

These examples explain the basics of Golang. There will be more examples from time to time.

if you like, feel free to add more Golang examples. Many thanks to all [contributors](https://github.com/SimonWaldherr/golang-examples/graphs/contributors).

## Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)

## Examples

The examples are divided into three levels of difficulty. The [Beginner](https://github.com/SimonWaldherr/golang-examples#beginner) section contains very easy examples, starting with **Hello World** but also containing a few easy algorithms. The [Advanced](https://github.com/SimonWaldherr/golang-examples#advanced) section uses more complicated features of Golang. Finally, the [Expert](https://github.com/SimonWaldherr/golang-examples#expert) section contains applications like telnet-clients or http-server (even with SSL).
If you want even more Golang examples, you can take a look at my other go repositories at GitHub:

* [golang-benchmarks](https://github.com/SimonWaldherr/golang-benchmarks) shows how to benchmark the execution time of Golang functions
* [GolangSortingVisualization](https://github.com/SimonWaldherr/GolangSortingVisualization) visualizes various sorting algorithms on the terminal or as gif
* [golang-minigames](https://github.com/SimonWaldherr/golang-minigames) currently only contains a snake clone
* [bbmandelbrot.go](https://github.com/SimonWaldherr/bbmandelbrot.go) calculates a [Mandelbrot Fractal](https://en.wikipedia.org/wiki/Mandelbrot_set) and saves it as PNG
* [golibs](https://github.com/SimonWaldherr/golibs) contains various Go packages (e.g. math, converter, stack, cli, ...)
* [fsagent](https://github.com/SimonWaldherr/fsagent) watch a folder for new or modified files and do something
* [cgol.go](https://github.com/SimonWaldherr/cgol.go) is [Conway's](https://en.wikipedia.org/wiki/John_Horton_Conway) [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in [Golang](https://en.wikipedia.org/wiki/Go_(programming_language))
* [micromarkdownGo](https://github.com/SimonWaldherr/micromarkdownGo) converts markdown to html (via regular expression)
* [wikiGo](https://github.com/SimonWaldherr/wikiGo) is a wiki software in Go
* [zplgfa](https://github.com/SimonWaldherr/zplgfa) is an image converter to print pictures on zpl compatible labels
* [...](https://github.com/search?utf8=✓&q=user%3Asimonwaldherr&type=Repositories&ref=advsearch&l=Go)

All of them are published as free and open source software.

If all of this is even not enough for you, you can take a look at the following websites:

* [tour.golang.org](https://tour.golang.org/)
* [Go by example](https://gobyexample.com/)
* [Golang Book](http://www.golang-book.com/)
* [Go-Learn](https://github.com/skippednote/Go-Learn)

### Beginner

To execute a **Golang** program, write ```go run``` at the cli followed by the name of the file.  
You also can convert the file to a binary executable program by the command ```go build```.  
If you know ```#!```, also known as [Shebang](https://en.wikipedia.org/wiki/Shebang_(Unix)), there is an equivalent for go: ```//usr/bin/env go run $0 $@ ; exit```  

Print Hello World with comments ([Golang Playground](https://play.golang.org/p/PiUVBrRB9AR))

```Shell
go run HelloWorld.go
```

Print Hello World with comments (shebang version)

```Shell
./HelloWorldShebang.go
```

Declare variables and print them ([Golang Playground](https://play.golang.org/p/O3_FWH2IQ75))

```Shell
go run var.go
```

Various ways (and styles) to print variables ([Golang Playground](https://play.golang.org/p/QH05DN_CeJU))

```Shell
go run printf.go
```

If statement in Golang ([Golang Playground](https://play.golang.org/p/rIpps0zUl1N))

```Shell
go run if.go Hello
```

Declare array and print its items ([Golang Playground](https://play.golang.org/p/jqseOd76Dqk))

```Shell
go run array.go
```

Declare your own functions ([Golang Playground](https://play.golang.org/p/9L0-F76gK0D))

```Shell
go run function.go
```

Do something multiple times ([Golang Playground](https://play.golang.org/p/5G3Ek89eKCH))

```Shell
go run for.go
```

Read via cli provided input data ([Golang Playground](https://play.golang.org/p/oUZ97D0Kg-_O))

```Shell
go run args.go string string2
```

Read via cli provided input data ([Golang Playground](https://play.golang.org/p/4IjMGLZBmue))

```Shell
go run input.go
```

Or scan for it ([Golang Playground](https://play.golang.org/p/rZ6clB9Z9Zu))

```Shell
go run scan.go
```

Read named argument input data ([Golang Playground](https://play.golang.org/p/i7BXecoK_wZ))

```Shell
go run flag.go
```

Return the *working directory* ([Golang Playground](https://play.golang.org/p/Ijx04bm8r8s))

```Shell
go run dir.go
```

Return the current time/date in various formats ([Golang Playground](https://play.golang.org/p/tPbn_CbHYYw))

```Shell
go run time.go
```

Return pseudo random integer values ([Golang Playground](https://play.golang.org/p/PapALcxQkpN))

```Shell
go run random.go
```

Concat strings in two different ways ([Golang Playground](https://play.golang.org/p/sBD_-peBzm5))

```Shell
go run cat.go
```

Modulo operation finds the remainder of division ([Golang Playground](https://play.golang.org/p/mYcueUHX1XL))

```Shell
go run modulo.go
```

Split a string by another string and make an array from the result ([Golang Playground](https://play.golang.org/p/VYP9tVMcbPW))

```Shell
go run split.go
```

An example implementation of the Ackermann function ([Golang Playground](https://play.golang.org/p/OjNgscvZhZT))

```Shell
go run ackermann.go
```

An example implementation of the Euclidean algorithm ([Golang Playground](https://play.golang.org/p/b1HWeUyk4Uf))

```Shell
go run euklid.go
```

Submit a function as argument ([Golang Playground](https://play.golang.org/p/fy0_S3J29_2))

```Shell
go run functioncallback.go
```

A function returned by a function ([Golang Playground](https://play.golang.org/p/earjKURMsPp))

```Shell
go run functionclosure.go
```

A function with an unknown amount of inputs (variadic function) ([Golang Playground](https://play.golang.org/p/s4Fvj8voh3Y))

```Shell
go run functionvariadic.go
```

Empty interface as argument (You Don't Know Type) ([Golang Playground](https://play.golang.org/p/D7OPshRATt_e))

```Shell
go run interface.go
```

Execute Shell/Bash commands and print its output values ([Golang Playground](https://play.golang.org/p/4Z-wLSkPJll))

```Shell
go run shell.go
```

Make structs (objects) which have functions ([Golang Playground](https://play.golang.org/p/t-082xlTu2t))

```Shell
go run oop.go
```

Dependency injection for easier testing

```Shell
cd beginner/di
go test
```

Hashing (md5, sha) in go ([Golang Playground](https://play.golang.org/p/fB2Y2MV7zt3))

```Shell
go run hashing.go
```

### Advanced

Benchmarking example (using JSON marshal and unmarshal for the sample) ([Golang Playground](https://play.golang.org/p/80Tcnkb301J)) 
From the root directory (`$GOPATH/github.com/SimonWaldherr/golang-examples`), run this command:

```Shell
go test -bench=. -benchmem advanced/json_bench/main_test.go
```

Make pipeable unix applications with os.Stdin ([Golang Playground](https://play.golang.org/p/NqrUOfBmJtt))

```Shell
go run pipe.go
```

AES-GCM encryption example ([Golang Playground](https://play.golang.org/p/ujfs6s5JZ-P))

```Shell
go run aesgcm.go
```

Bcrypt hashing example ([Golang Playground](https://play.golang.org/p/9R7oS56Od6H)) 
Please install package golang.org/x/crypto/bcrypt before run this file by running `go get golang.org/x/crypto/bcrypt`

```Shell
go run bcrypt.go
```

Search element is exist in arrays or not ([Golang Playground](https://play.golang.org/p/1gVa9Jgk6vg))

```Shell
go run in_array.go
```

Calculate triangles ([Golang Playground](https://play.golang.org/p/l8ehuAWZitv))

```Shell
go run pythagoras.go (float|?) (float|?) (float|?)
```

Read from stdin (but don't wait for the enter key)

```Shell
go run getchar.go
```

Wait and sleep ([Golang Playground](https://play.golang.org/p/qGec1g7rTHC))

```Shell
go run wait.go
```

Last in - first out - example (Pop and push in Golang) ([Golang Playground](https://play.golang.org/p/TekltztwUfE))

```Shell
go run lifo.go
```

Split a string via regular expression and make an array from the result ([Golang Playground](https://play.golang.org/p/sWFDPMyF-wD))

```Shell
go run regex.go
```

More advanced regex (with time and dates) ([Golang Playground](https://play.golang.org/p/u1SdhFEwRch))

```Shell
go run regex2.go
```

Use my [golibs regex package](https://github.com/SimonWaldherr/golibs#regex-----) and have fun ([Golang Playground](https://play.golang.org/p/1RxtOxL0nQo))

```Shell
go run regex3.go
```

Calculate and print the fibonacci numbers ([Golang Playground](https://play.golang.org/p/BbjQ_ohw0m1))

```Shell
go run fibonacci.go
```

Calculate and print the requested (32th) prime number ([Golang Playground](https://play.golang.org/p/fB25PQlVWu4))

```Shell
go run prime.go 32
```

Do things with numbers, strings and switch-cases ([Golang Playground](https://play.golang.org/p/MvKDvR_wzlQ))

```Shell
go run numbers.go
```

Use a template to create and fill documents (this example uses [LaTeX](https://www.latex-project.org)) ([Golang Playground](https://play.golang.org/p/riy6SU21alH))

```Shell
go run template.go
pdflatex -interaction=nonstopmode template_latex.tex
```

Start a ticker (do things periodically) 

```Shell
go run ticker.go
```

Do something in case of a timeout ([Golang Playground](https://play.golang.org/p/Qtu62LaC3-q))

```Shell
go run timeout.go
```

Convert go object to json string ([Golang Playground](https://play.golang.org/p/s1onU7jpP91))

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

Parse CSV ([Golang Playground](https://play.golang.org/p/plVmXW8TB32))

```Shell
go run csv.go
```

Convert CSV to a Markdown table ([Golang Playground](https://play.golang.org/p/HTobb_U9JQt))

```Shell
go run csv2md.go
```

Parse a XML string into a Struct with undefined Fields ([Golang Playground](https://play.golang.org/p/6LuNPcaeagS))

```Shell
go run xml.go
```

Run a self killing app

```Shell
go run suicide.go
```

GoCV : hello video

```Shell
go run hello_video.go
```

GoCV : face detection

```Shell
go run face_detect.go 0 model/haarcascade_frontalface_default.xml
```

Run the example for generic ([Golang Playground](https://go.dev/play/p/586EEI0ZAEe))

```Shell
go run generic.go
```

### Expert

Calculate π with go (leibniz, euler and prime are running until you stop it via CTRL+C)

```Shell
go run pi2go.go leibniz
go run pi2go.go euler
go run pi2go.go prime
```

Calculate π with go - same as above - but with live output (based on [gcurses](https://godoc.org/github.com/SimonWaldherr/golibs/gcurses))

```Shell
go run pi2go-live.go leibniz
go run pi2go-live.go euler
go run pi2go-live.go prime
```


List files in working directory

```Shell
go run explorer.go
```

run assemply code from golang

```Shell
go run assembly.go
```

run C code from golang 

```Shell
go run cgo.go
```

generate Go code with golang templates

```Shell
go run codegen.go
```

Convert from rgb to hsl ([Golang Playground](https://play.golang.org/p/UuX27PhA0Zx))

```Shell
go run color.go
```

Telnet with Golang

```Shell
go run telnet.go
```

The smallest Golang http server

```Shell
go run httpd.go
```

Secure Golang http server

```Shell
go run httpsd.go
```

The smallest Golang http proxy

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

Generate an image with three colored circles (with intersection)

```Shell
go run image2.go
```

Generate an image representing the Mandelbrot fractal

```Shell
go run image3.go
```

Sql (sqlite) Golang example  
maybe you also wanna take a look at my [sql-examples](https://github.com/SimonWaldherr/sql-examples)-project

```Shell
go run sqlite.go insert test
go run sqlite.go select
```

Public-key/asymmetric cryptography signing and validating

```Shell
go run ppk-crypto.go
```

Command Line Arguments Golang Example
We can get argument values though command line by specifying the operator '-' with the name of the argument and the value to be set. E.g. -env=qa

```Shell
go run command_line_arguments.go
go run command_line_arguments.go -env=qa -consumer=true
```

Cron Golang Example
We can trigger a function at a particular time through cron 

```Shell
go run cron.go
```

Map Golang Example
Hash Map standard functions in golang 

```Shell
go run map.go
```

## Compile

One great aspect of Golang is, that you can start go applications via ```go run name.go```, but also compile it to an executable with ```go build name.go```. After that you can start the compiled version which starts much faster.
If you start fibonacci.go and the compiled version you will notice, that the last line which contains the execution time doesn't differ much, but if you start it with ```time ./fibonacci 32``` and ```time go run ./fibonacci.go 32``` you will see the difference.

## License

Copyright © 2022 Simon Waldherr
Dual-licensed. See the [LICENSE](https://github.com/SimonWaldherr/golang-examples/blob/master/LICENSE) file for details.

## Support me

if you like what i do
feel free to support me

you can do so by:

* [donate via PayPal](https://www.paypal.me/SimonWaldherr "Donate to this project via PayPal.me") or [liberaPay](https://liberapay.com/SimonWaldherr/donate "Donate using Liberapay")
* buy me a beer or [Club-Mate](https://en.wikipedia.org/wiki/Club-Mate#Hacker_culture) at a conference
* give me a job where I can work on open source projects (please don't contact me via LinkedIn - please send an eMail or [contact me via twitter](http://twitter.com/SimonWaldherr) instead)
