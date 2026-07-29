# Golang-Examples

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21677192.svg)](https://doi.org/10.5281/zenodo.21677192)
[![Modern Go examples](https://github.com/SimonWaldherr/golang-examples/actions/workflows/modern-go.yml/badge.svg)](https://github.com/SimonWaldherr/golang-examples/actions/workflows/modern-go.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Small, focused programs for learning Go—from `Hello, World!` to concurrency,
WebAssembly, runtime tracing, and modern language features. Most examples are
standalone programs and can be run independently.

Use the [online editor](https://simonwaldherr.github.io/golang-examples/) to
edit and run examples in the browser, or play the WebAssembly-based
[Qix game](https://simonwaldherr.github.io/golang-examples/non-std-lib/ebiten.html).

## Quick start

The module version in [`go.mod`](go.mod) is the source of truth for the
required Go toolchain.

```shell
git clone https://github.com/SimonWaldherr/golang-examples.git
cd golang-examples
go run ./beginner/HelloWorld.go
go run ./advanced/iterators
./scripts/check-modern.sh
```

## Repository map

| Directory | What you will find |
| --- | --- |
| [`beginner/`](beginner/) | Syntax, types, control flow, files, and small algorithms |
| [`advanced/`](advanced/) | Generics, iterators, concurrency, encoding, tests, and benchmarks |
| [`expert/`](expert/) | Servers, tracing, assembly, CGO, cryptography, and image generation |
| [`non-std-lib/`](non-std-lib/) | Examples that depend on third-party packages or external services |
| [`tinygo/`](tinygo/) | TinyGo and microcontroller examples |

## Modern Go highlights

These examples intentionally use recent stable Go features:

| Go version | Feature | Run it |
| --- | --- | --- |
| 1.23 | Range-over-function iterators plus `iter`, `maps`, and `slices` | `go run ./advanced/iterators` |
| 1.24 | Generic type aliases | `go run ./advanced/generic-alias` |
| 1.24 | Benchmarks with `testing.B.Loop` | `go test ./advanced/benchmark-loop -bench .` |
| 1.25 | `sync.WaitGroup.Go` | `go run ./advanced/waitgroup-go` |
| 1.25 | Deterministic concurrent tests with `testing/synctest` | `go test ./advanced/synctest` |
| 1.25 | Runtime trace flight recorder | `go run ./expert/flight-recorder` |
| 1.26 | Initialized pointers with `new(expression)` | `go run ./beginner/initialized-pointer` |
| 1.21 | Structured JSON logging with `log/slog` | `go run ./advanced/structured-logging` |

The flight-recorder example writes `flight.trace`; inspect it with
`go tool trace flight.trace`.

## Featured project: nanoGo

[nanoGo](https://github.com/SimonWaldherr/nanoGo) is a minimalist interpreter
for a supported subset of Go. It evaluates source code dynamically in a CLI,
REPL, embedded host, or browser playground; when used in the browser, the
interpreter itself is compiled to WebAssembly. That makes nanoGo a strong fit
for interactive tutorials, editable documentation, controlled snippets, and
browser-based Go experiments. Try the
[nanoGo playground](https://simonwaldherr.github.io/nanoGo/) or embed it in a
web page.

### nanoGo, TinyGo, and GopherJS solve different problems

| | [nanoGo](https://github.com/SimonWaldherr/nanoGo) | [TinyGo](https://tinygo.org/) | [GopherJS](https://github.com/gopherjs/gopherjs) |
| --- | --- | --- | --- |
| Execution model | Interprets supported Go source at runtime | Compiles Go programs ahead of time | Compiles Go programs ahead of time to JavaScript |
| Browser artifact | The interpreter runs in WASM and evaluates guest source dynamically | The application itself can be compiled to WASM | Pure JavaScript generated from the application |
| Best suited to | Playgrounds, REPLs, live examples, controlled embedded scripting, and teaching | Microcontrollers, embedded applications, and deployable WASM programs | Browser front ends and JavaScript-based web integrations |
| Go compatibility | Deliberately supported language and library subset | A compiler with its own documented Go compatibility differences | Broad Go support with documented browser and JavaScript-runtime constraints |
| Host control | Optional capabilities and cooperative resource limits can restrict guest source | The compiled program runs for its selected target; TinyGo is not an interpreter sandbox | The generated program executes with normal browser JavaScript capabilities |

nanoGo is therefore not a smaller replacement for TinyGo or GopherJS. Choose
**nanoGo** when source must be edited or evaluated at runtime; choose
**TinyGo** when you want to compile and deploy an application, for example to
a Raspberry Pi Pico; choose **GopherJS** when compiled JavaScript is the right
browser target.

### The existing live editor uses GopherJS

The [GitHub Pages editor](https://simonwaldherr.github.io/golang-examples/)
uses its [`go2js`](https://github.com/live-codes/go2js) integration, which is
based on GopherJS. Pressing <kbd>F5</kbd> formats the editor contents,
compiles the Go program to JavaScript, and evaluates that JavaScript in the
page. It is a compile-and-run workflow, not a nanoGo interpreter session.

## Related projects

The following curated list contains public, non-fork repositories from
[SimonWaldherr](https://github.com/SimonWaldherr) that complement this
collection.

### Go learning and reusable tools

- [golang-benchmarks](https://github.com/SimonWaldherr/golang-benchmarks) — examples for measuring Go code
- [GolangSortingVisualization](https://github.com/SimonWaldherr/GolangSortingVisualization) — visualized sorting algorithms
- [golibs](https://github.com/SimonWaldherr/golibs) — general-purpose Go packages
- [gotools](https://github.com/SimonWaldherr/gotools) — a collection of small Go tools
- [GoRealtimeWeb](https://github.com/SimonWaldherr/GoRealtimeWeb) — real-time web application examples
- [mdExec](https://github.com/SimonWaldherr/mdExec) — executes code blocks in Markdown files

### Runtimes, data, and AI

- **[nanoGo](https://github.com/SimonWaldherr/nanoGo)** — a Go-subset interpreter for native hosts and WebAssembly, with a playground, REPL, CLI, and embeddable host API
- [tinySQL](https://github.com/SimonWaldherr/tinySQL) — an educational SQL engine written in pure Go
- [tinyRAG](https://github.com/SimonWaldherr/tinyRAG) — a lightweight retrieval-augmented generation system
- [smallR](https://github.com/SimonWaldherr/smallR) — a small R-like environment written in Go
- [DataDock](https://github.com/SimonWaldherr/DataDock) — a server-side database web interface

### Graphics, games, and hardware

- [golang-minigames](https://github.com/SimonWaldherr/golang-minigames) — small games written in Go
- [bbmandelbrotGo](https://github.com/SimonWaldherr/bbmandelbrotGo) — Mandelbrot image generation
- [FluidSimASCII](https://github.com/SimonWaldherr/FluidSimASCII) — an ASCII fluid simulator
- [vango](https://github.com/SimonWaldherr/vango) — image-manipulation effects
- [rp2040-examples](https://github.com/SimonWaldherr/rp2040-examples) and [rpi-examples](https://github.com/SimonWaldherr/rpi-examples) — Raspberry Pi and RP2040 examples
- [RGB-LED-Matrix](https://github.com/SimonWaldherr/RGB-LED-Matrix) and [pico75player](https://github.com/SimonWaldherr/pico75player) — LED-matrix projects

### Other example collections

- [sql-examples](https://github.com/SimonWaldherr/sql-examples)
- [openscad-examples](https://github.com/SimonWaldherr/openscad-examples)
- [zig-examples](https://github.com/SimonWaldherr/zig-examples)
- [cobol-examples](https://github.com/SimonWaldherr/cobol-examples)

All are published as free and open-source software. Browse the complete
[Go repository search](https://github.com/search?q=user%3ASimonWaldherr+language%3AGo&type=repositories)
for more.

## Install Go

- macOS with Homebrew: `brew install go`
- Debian/Ubuntu: `sudo apt install golang-go`
- Other systems: follow the official [Go installation guide](https://go.dev/doc/install)

## Example catalog

The examples are divided into beginner, advanced, expert, third-party, and
TinyGo sections. Commands in the catalog below are shown relative to the
corresponding directory unless they include a directory prefix.

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

Error handling – creating, returning, wrapping and inspecting errors

```Shell
go run error.go
```

Switch statement – expression switch, condition switch, type switch and fallthrough

```Shell
go run switch.go
```

Type conversions – numeric casts, string/[]byte/[]rune and strconv helpers

```Shell
go run typeconv.go
```

### Advanced

Benchmarking example (using JSON marshal and unmarshal for the sample) ([Golang Playground](https://play.golang.org/p/80Tcnkb301J)) 
From the root directory (`$GOPATH/github.com/SimonWaldherr/golang-examples`), run this command:

```Shell
go test -bench=. -benchmem advanced/json_bench/main_test.go
```

Make pipe-able unix applications with os.Stdin ([Golang Playground](https://play.golang.org/p/NqrUOfBmJtt))

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

Protect shared state with sync.Mutex and sync.RWMutex

```Shell
go run mutex.go
```

Context cancellation, timeouts, deadlines, and value propagation

```Shell
go run context.go
```

Worker pool – distribute jobs across a fixed number of goroutines

```Shell
go run workerpool.go
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

run assembly code from golang

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

Token-bucket rate limiter – throttle request throughput with burst support

```Shell
go run ratelimiter.go
```

### TinyGo

You can even use Go on microcontrollers, the keyword here is [TinyGo](https://tinygo.org/), a go compiler specially developed for SBCs and MCUs.  
This is distinct from [nanoGo](https://github.com/SimonWaldherr/nanoGo): TinyGo
compiles programs for deployment, while nanoGo interprets a supported Go subset
dynamically for interactive and embedded-host use cases.
If you want to blink the LED of your Raspberry Pi Pico, try this: 

```Shell
tinygo build -o firmware.uf2 -target=pico ./tinygo/blink.go
```

and then upload it to the pico.

## Compile

One great aspect of Golang is, that you can start go applications via ```go run name.go```, but also compile it to an executable with ```go build name.go```. After that you can start the compiled version which starts much faster.
If you start fibonacci.go and the compiled version you will notice, that the last line which contains the execution time doesn't differ much, but if you start it with ```time ./fibonacci 32``` and ```time go run ./fibonacci.go 32``` you will see the difference.

## License

Copyright © 2026 Simon Waldherr
Dual-licensed. See the [LICENSE](https://github.com/SimonWaldherr/golang-examples/blob/master/LICENSE) file for details.
