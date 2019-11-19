# go-repr

[![GoDoc Reference][godoc_badge]][godoc]
[![Build Status][actions_badge]][actions]

String representations for Go values. 💬

## Installation

``go get github.com/hackebrot/go-repr``

## Usage

The primary feature of **go-repr** is a function for generating string
representations for Go values: ``repr.Repr()``. It resolves pointers to their
values and omits unexported struct fields as well as struct fields with nil
values.

### Example

```go
package main

import (
	"fmt"

	"github.com/hackebrot/go-repr/repr"
)

// Emoji holds info about an emoji character
type Emoji struct {
	Name     string
	Category string
	Char     string
	Keywords []string
}

func main() {
	fmt.Println(repr.Repr("hello world! 🌍"))
	fmt.Println(repr.Repr(1234))
	fmt.Println(repr.Repr(true))

	astronaut := Emoji{
		Name:     "woman_astronaut",
		Category: "people",
		Char:     "👩‍🚀",
		Keywords: []string{"space", "rocket", "woman", "human"},
	}

	fmt.Println(repr.Repr(&astronaut))
}
```

```text
$ go run example.go
"hello world! 🌍"
1234
true
main.Emoji{Name:"woman_astronaut", Category:"people", Char:"👩‍🚀", Keywords:["space" "rocket" "woman" "human"]}
```

## About

This project is inspired by ``github.Stringify()`` of
[google/github-go][go-github] and aims at making it easier to work with
structs that use pointer fields. Without ``repr.Repr()``, you would have to
check for ``nil`` for all of your struct's pointer fields before you can
derefence them one by one.

For more information about **go-repr** please check out [my blog][blog]. 📝

## Community

Contributions are welcome, and they are greatly appreciated! Every little bit
helps, and credit will always be given. Please check out this
[guide][contributing] to get started!

Please note that this project is released with a [Contributor Code of
Conduct][Code of Conduct]. By participating in this project you agree to
abide by its terms.

## License

Distributed under the terms of the [MIT License][MIT], **go-repr** is free
and open source software.

[blog]: https://raphael.codes/blog/string-representations-for-go-values/
[Code of Conduct]: /CODE_OF_CONDUCT.md
[contributing]: /CONTRIBUTING.md
[go-github]: https://github.com/google/go-github
[godoc_badge]: https://img.shields.io/badge/go-documentation-blue.svg?style=flat
[godoc]: https://godoc.org/github.com/hackebrot/go-repr/repr (See GoDoc Reference)
[MIT]: /LICENSE
[actions_badge]: https://github.com/hackebrot/go-repr/workflows/test/badge.svg
[actions]: https://github.com/hackebrot/go-repr/actions (See Build Status on GitHub Actions)
