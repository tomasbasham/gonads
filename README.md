# gonads

In [functional
programming](https://en.wikipedia.org/wiki/Functional_programming) a monad is
a structure that encapsulates a function, and its return value(s), in a type
providing more deterministic computational behaviours, and mitigating
side-effects. Mondas may be used to reduce boilerplate code needed for common
operations, such as dealing with optional or undefined values safely.

This module implements some well-known monadic structures that can be used in
Go programs to introduce functional paradigms.

## Prerequisites

You will need the following things properly installed on your computer.

- [Go](https://golang.org/): any one of the **three latest major**
  [releases](https://golang.org/doc/devel/release.html)

## Installation

With [Go module](https://github.com/golang/go/wiki/Modules) support (Go 1.11+),
simply add the following import

```go
import "github.com/tomasbasham/gonads"
```

to your code, and then `go [build|run|test]` will automatically fetch the
necessary dependencies.

Otherwise, to install the `gonads` module, run the following command:

```bash
$ go get -u github.com/tomasbasham/gonads
```

## License

This project is licensed under the [MIT License](LICENSE.md).
