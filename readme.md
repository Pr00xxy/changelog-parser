# Golang changelog parser

Small self contained library for reading changelog following the [keep a changelog](https://keepachangelog.com/en/1.0.0/) format.  
This is released in the hopes that it will be useful

## Installation

    go get github.com/prooxxy/changelog-parser

## Usage

The New() api takes an io.Buffer and Parse returns the Changelog object

```go
package main

import (
    parser "github.com/prooxx/changelog-parser"
)

func main() {
    p, _ := parser.New(<io.Buffer>)

    cl := p.Parse()

    versions := cl.Versions
}

```

## Roadmap
Non currently.
Needs refactoring

## License
changelog-parser is released under the [MIT License](https://github.com/victorspringer/http-cache/blob/master/LICENSE)