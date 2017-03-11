# go-repr

[![GoDoc Reference][godoc_badge]][godoc]
[![Build Status][travis_badge]][travis]

String representations of go objects

## Installation

``go get github.com/hackebrot/go-repr``

## Usage

Import **go-repr** and then use ``repr.Repr()`` to create a string
representation for an object. It resolves pointers to their values and omits
unexported struct fields as well as struct fields with nil values.

```go
repr.Repr("hello world!")

repr.Repr(1234)

type Gopher struct {
    Hair *string
    Eyes *string
}

h := "rainbow"
e := "Goofy Eyes"
g := &Gopher{Hair: &h, Eyes: &e}

repr.Repr(g)
```

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/hackebrot/go-repr/repr"
)

// Maintainer of a Project
type Maintainer struct {
	Name *string
}

// Release represents a version of a Project
type Release struct {
	PublishedAt time.Time
	Version     string
}

// Project represents a OSS project
type Project struct {
	Name        *string
	Description *string
	Maintainers []*Maintainer
	Forks       *int
	url         *string
	Platform    map[string]bool
	Releases    []*Release
}

func main() {
	aName := "Amelia"
	amelia := &Maintainer{Name: &aName}

	tName := "Tony"
	tony := &Maintainer{Name: &tName}

	pName := "Chew"
	pForks := 550

	p := &Project{
		Name:        &pName,
		Maintainers: []*Maintainer{amelia, tony},
		Forks:       &pForks,
		Platform:    map[string]bool{"linux": true, "windows": true, "osx": true},
		Releases: []*Release{
			&Release{
				PublishedAt: time.Date(2016, 01, 01, 01, 04, 10, 0, time.UTC),
				Version:     "0.1.0",
			},
			&Release{
				PublishedAt: time.Date(2017, 01, 02, 15, 04, 05, 0, time.UTC),
				Version:     "1.0.0",
			},
		},
	}

	fmt.Printf("%v\n", repr.Repr(p))
}
```

Above code will generate the following output:

```text
main.Project{Name:"Chew", Maintainers:[main.Maintainer{Name:"Amelia"}
main.Maintainer{Name:"Tony"}], Forks:550, Platform:map["windows":true "osx":true 
"linux":true], Releases:[main.Release{PublishedAt:time.Time{2016-01-01 01:04:10 
+0000 UTC}, Version:"0.1.0"} main.Release{PublishedAt:time.Time{2017-01-02 
15:04:05 +0000 UTC}, Version:"1.0.0"}]}
```


## About

**go-repr** is inspired by ``github.Stringify()`` of
[google/github-go][go-github] and aims at making it easier to work with
structs that use pointer fields. Without ``github.Stringify()`` or
``repr.Repr()`` respectively, you would have to check for ``nil`` for all of
your struct's pointer fields before you can derefence them one by one.

**go-repr** allows you to *debug print* your objects with a single line of
code.

## License

Distributed under the terms of the [MIT License][MIT], **go-repr** is
free and open source software.


## Contributing

Contributions are welcome, and they are greatly appreciated! Every
little bit helps, and credit will always be given.

Please check out this [guide][contributing] to get started!


## Code of Conduct

Please note that this project is released with a
[Contributor Code of Conduct][Code of Conduct].

By participating in this project you agree to abide by its terms.


[Code of Conduct]: CODE_OF_CONDUCT.md
[contributing]: CONTRIBUTING.md
[go-github]: https://github.com/google/go-github
[godoc_badge]: https://img.shields.io/badge/go-documentation-blue.svg?style=flat
[godoc]: https://godoc.org/github.com/hackebrot/go-repr (See GoDoc Reference)
[MIT]: LICENSE
[travis_badge]: https://img.shields.io/travis/hackebrot/go-repr.svg?style=flat
[travis]: https://travis-ci.org/hackebrot/go-repr (See Build Status on Travis CI)
