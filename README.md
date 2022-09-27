# ci-detective
Simple utility to detect whether you are running in a CI server/environment

[![Latest Release][github-release-badge]][github-release-url]
[![GoDoc][godoc-badge]][godoc-url]
[![License Badge][license-badge]][license-url]  

[![Linux CI Badge][linux-ci-badge]][linux-ci-url]
[![Mac CI Badge][mac-ci-badge]][mac-ci-url]
[![Windows CI Badge][windows-ci-badge]][windows-ci-url]  

[![Go Report Card][go-report-card-badge]][go-report-card-url]
[![Test Results Badge][tests-badge]][sonar-tests-url]
[![Codecov Badge][codecov-badge]][codecov-url]
[![Sonar Quality Gate Badge][sonar-quality-gate-badge]][sonar-url]  

<br />

## Installation
`ci-detective` is available as both an executable/cli and a package/api.

Grab a [binary release][github-release-url] or use `go get`: 

```sh
go get -u github.com/swellaby/ci-detective
```

If you want to consume the package/api programmatically, just add the package import:

```go
import (
    "github.com/swellaby/ci-detective/cidetective"
)
```

## Usage
### CLI
To use `ci-detective` as a CLI simply run the command:

```sh
ci-detective
```

The CLI will return an exit code of Zero if it detects the current environment is a CI engine/server. Otherwise it will return a non-Zero exit code.

### Lib
The `cidetective` package exposes a single function: `IsCI` that returns a boolean indicating whether or not the current environment is a CI server/engine.

```go
package main

import (
    "fmt"
    "github.com/swellaby/ci-detective/cidetective"
)

func main() {
    if cidetective.IsCI() {
        fmt.Println("Running in a CI environment!")
    } else {
        fmt.Println("Nope, this is NOT a CI environment!")
    }
}
```


### Generator
Initially created by this [swell generator][parent-generator-url]!

### Supported CIs

- [Woodpecker CI][woodpecker-ci]

[github-release-badge]: https://img.shields.io/github/release/swellaby/ci-detective?style=flat-square
[github-release-url]: https://github.com/swellaby/ci-detective/releases/latest
[go-report-card-badge]: https://goreportcard.com/badge/github.com/swellaby/ci-detective?style=flat-square
[go-report-card-url]: https://goreportcard.com/report/github.com/swellaby/ci-detective
[godoc-badge]: https://img.shields.io/badge/godoc-reference-blue?style=flat-square
[godoc-url]: https://godoc.org/github.com/swellaby/ci-detective/cidetective
[license-url]: ./LICENSE
[license-badge]: https://img.shields.io/github/license/swellaby/ci-detective.svg?style=flat-square&color=blue
[linux-ci-badge]: https://img.shields.io/azure-devops/build/swellaby/opensource/134/master?style=flat-square&label=linux%20build
[linux-ci-url]: https://dev.azure.com/swellaby/OpenSource/_build/latest?definitionId=134
[mac-ci-badge]: https://img.shields.io/azure-devops/build/swellaby/opensource/131/master?style=flat-square&label=mac%20build
[mac-ci-url]: https://dev.azure.com/swellaby/OpenSource/_build/latest?definitionId=131
[windows-ci-badge]: https://img.shields.io/azure-devops/build/swellaby/opensource/133/master?style=flat-square&label=windows%20build
[windows-ci-url]: https://dev.azure.com/swellaby/OpenSource/_build/latest?definitionId=133
[codecov-badge]: https://img.shields.io/codecov/c/github/swellaby/ci-detective.svg?style=flat-square
[codecov-url]: https://codecov.io/gh/swellaby/ci-detective
[tests-badge]: https://img.shields.io/azure-devops/tests/swellaby/opensource/134?style=flat-square
[sonar-quality-gate-badge]: https://img.shields.io/sonar/quality_gate/swellaby:ci-detective?server=https%3A%2F%2Fsonarcloud.io&style=flat-square
[sonar-url]: https://sonarcloud.io/dashboard?id=swellaby%3Aci-detective
[sonar-tests-url]: https://sonarcloud.io/component_measures?id=swellaby%3Aci-detective&metric=tests
[parent-generator-url]: https://github.com/swellaby/generator-lets-go
[woodpecker-ci]: https://woodpecker-ci.org
