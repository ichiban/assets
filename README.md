# assets
Go resource embedding library which doesn't require `go generate` at all.

[![GoDoc](https://godoc.org/github.com/ichiban/assets?status.svg)](https://godoc.org/github.com/ichiban/assets)
[![Build Status](https://travis-ci.org/ichiban/assets.svg?branch=master)](https://travis-ci.org/ichiban/assets)

## Overview

[Resource Embedding in Go](https://awesome-go.com/#resource-embedding) is cumbersome especially when it involves custom commands and `go generate`.
`assets` is another resource embedding library which doesn't require any of them.
It extracts a zip file bundled with your binary and provides a file path to the contents.

## Features

- no custom commands
- no `go generate`
- only requires `cat` and `zip` when `go build`
- uses local files when `go run` or `go test`
- easily extendable interface

## Getting Started

### Prerequisites

We assume you have a Go project and a directory named `assets` for resources in the project root.
This directory name can be easily configured but for the sake of explanation we call the directory `assets`.

```
$ tree .
.
├── Makefile
├── assets
│   └── templates
│       └── hello.tmpl
└── main.go

2 directories, 3 files
``` 

### Install

First go get the library.

```shell
$ go get github.com/ichiban/assets
```

Then import it.

```
import "github.com/ichiban/assets"
```

### Usage

In Go files, we can get the locator which points to `assets` directory by calling `assets.New()`.
Note that we'll need to close it in the end.

```go
	l, err := assets.New()
	if err != nil {
		log.Fatalf("assets.New() failed: %v", err)
	}
	defer l.Close()
	
	log.Printf("assets: %s", l.Path)
```

When we build it and bundle with a resource zip file, the locator will point to the contents of the zip file which is extracted into a temporary directory.

### Build

To start with, we need to build the binary as usual.

```
$ mkdir -p bin 
$ go build -o bin/hello
```

Then, zip resources into a zip file.
We need to enter `assets` directory to make a proper zip file.

```
$ mkdir -p zip
$ cd assets
$ zip -r ../zip/assets.zip .
  adding: templates/ (stored 0%)
  adding: templates/hello.tmpl (stored 0%)
$ cd ..
```

A proper zip file looks like this:

```
$ unzip -l zip/assets.zip 
Archive:  zip/assets.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  11-05-2018 22:31   templates/
       14  11-05-2018 21:31   templates/hello.tmpl
---------                     -------
       14                     2 files
```

Finally, we can bundle resources by `cat` the binary and the zip file.
Prepending an executable makes the zip offset values off by the size of the executable.
We can fix the offsets by `zip -A` (`zip --adjust-sfx`).

```
$ cat bin/hello zip/assets.zip > bin/hello-bundled
$ zip -A bin/hello-bundled 
Zip entry offsets appear off by 3345496 bytes - correcting...
$ chmod +x bin/hello-bundled
```

Interestingly, the bundled binary is an executable and also a zip file.

```
$ unzip -l bin/hello-bundled 
Archive:  bin/hello-bundled
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  11-05-2018 22:31   templates/
       14  11-05-2018 21:31   templates/hello.tmpl
---------                     -------
       14                     2 files
```

```
$ bin/hello-bundled 
2018/11/30 16:22:13 assets: /var/folders/xt/6z9sk1dx1d734ltxxst16_h00000gn/T/hello-bundled882816570
Hello, World!
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details. 

## Acknowledgments

This project is inspired by [Zgok](https://github.com/srtkkou/zgok).