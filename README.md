# go-handcash-connect
> The unofficial golang implementation for the [HandCash Connect SDK](https://handcash.github.io/handcash-connect-sdk-js-beta-docs/#/)

[![Release](https://img.shields.io/github/release-pre/tonicpow/go-handcash-connect.svg?logo=github&style=flat&v=1)](https://github.com/tonicpow/go-handcash-connect/releases)
[![Build Status](https://travis-ci.com/tonicpow/go-handcash-connect.svg?branch=master&v=2)](https://travis-ci.com/tonicpow/go-handcash-connect)
[![Report](https://goreportcard.com/badge/github.com/tonicpow/go-handcash-connect?style=flat&v=2)](https://goreportcard.com/report/github.com/tonicpow/go-handcash-connect)
[![codecov](https://codecov.io/gh/tonicpow/go-handcash-connect/branch/master/graph/badge.svg)](https://codecov.io/gh/tonicpow/go-handcash-connect)
[![Go](https://img.shields.io/github/go-mod/go-version/tonicpow/go-handcash-connect)](https://golang.org/)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-handcash-connect** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/tonicpow/go-handcash-connect
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/tonicpow/go-handcash-connect)

[![GoDoc](https://godoc.org/github.com/tonicpow/go-handcash-connect?status.svg&style=flat)](https://pkg.go.dev/github.com/tonicpow/go-handcash-connect)

### Features
- [Client](client.go) is completely configurable
- Current coverage for the [HandCash Connect SDK](https://handcash.github.io/handcash-connect-sdk-js-beta-docs/#/)
    - [x] GetProfile
    - [ ] GetBalance
    - [ ] GetPayRequest


<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
clean                  Remove previous builds and any test cache data
clean-mods             Remove all the Go mod cache
coverage               Shows the test coverage
godocs                 Sync the latest tag with GoDocs
help                   Show this help message
install                Install the application
install-go             Install the application (Using Native Go)
lint                   Run the Go lint application
release                Full production release (creates release in Github)
release                Runs common.release then runs godocs
release-snap           Test the full release (build binaries)
release-test           Full production test release (everything except deploy)
replace-version        Replaces the version in HTML/JS (pre-deploy)
tag                    Generate a new tag and push (tag version=0.0.0)
tag-remove             Remove a tag if found (tag-remove version=0.0.0)
tag-update             Update an existing tag to current commit (tag-update version=0.0.0)
test                   Runs vet, lint and ALL tests
test-short             Runs vet, lint and tests (excludes integration tests)
test-travis            Runs all tests via Travis (also exports coverage)
test-travis-short      Runs unit tests via Travis (also exports coverage)
uninstall              Uninstall the application (and remove files)
vet                    Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](handcash_test.go) run via [Travis CI](https://travis-ci.org/tonicpow/go-handcash-connect) and uses [Go version 1.15.x](https://golang.org/doc/go1.15). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go [benchmarks](handcash_test.go):
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
View the [examples](handcash_test.go)

Basic implementation:
```go
package main

import (
    "github.com/tonicpow/go-handcash-connect"
)

func main() {

   // todo: add a basic example
}
```
 
<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) | [<img src="https://github.com/rohenaz.png" height="50" alt="Satchmo" />](https://github.com/rohenaz) |
|:---:|:---:|
| [MrZ](https://github.com/mrz1836) | [Satchmo](https://github.com/rohenaz) |
              
<br/>

## Contributing
View the [contributing guidelines](CONTRIBUTING.md) and please follow the [code of conduct](CODE_OF_CONDUCT.md).

### Credits

[HandCash](https://handcash.io) for their hard work on the HandCash Connect SDK

<br/>

## License

![License](https://img.shields.io/github/license/tonicpow/go-handcash-connect.svg?style=flat)