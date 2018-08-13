The Unofficial [Nomics](https://nomics.com/) API client for [Go](https://golang.org/).

## Documentation

[https://godoc.org/github.com/kilpatty/nomics-go/nomics](https://godoc.org/github.com/kilpatty/nomics-go/nomics)

## Install

```bash
go get -u github.com/kilpatty/nomics-go/nomics
```

## Getting started

```go
package main

import (
	"fmt"
	"log"

	"github.com/kilpatty/nomics-go/nomics"
)

func main() {

	config := nomics.New("Your-API-Key")

	markets, err := nomics.Markets()
	if err != nil {
		log.Fatal(err)
	}

	for _, market := range markets {
		fmt.Println(market.Market, market.Quote)
	}
}

```

## Examples

Check out the [`./example`](./example) directory and documentation.

## License

MIT
