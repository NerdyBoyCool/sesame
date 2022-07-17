## Sesame

[![Go Reference](https://pkg.go.dev/badge/github.com/NerdyBoyCool/sesame.svg)](https://pkg.go.dev/github.com/NerdyBoyCool/sesame)

This package provides a client to the Sesame API.

This package API Client to lock the keys of Sesame API using signatures encrypted with the AES-CMAC method.

For more information about the API, please see the official documentation.

https://doc.candyhouse.co/ja/SesameAPI

## Example
Toggle Sesame From Script

```go

package main

import (
	"context"
	"fmt"
	"github.com/NerdyBoyCool/sesame"
)

func main() {
	cli := sesame.NewClient("your sesame api key", "sesame secret key for aes-cmac", "Sesame UUID assigned to each Sesame(Sesame UUID)")
	ctx := context.Background()
	err := cli.Toggle(ctx, "From API")
	if err != nil {
		fmt.Println(err)
	}
}

```

Get current sesame's information.

```go

package main

import (
	"context"
	"fmt"
	"github.com/NerdyBoyCool/sesame"
)

func main() {
	cli := sesame.NewClient("your sesame api key", "sesame secret key for aes-cmac", "Sesame UUID assigned to each Sesame(Sesame UUID)")
	ctx := context.Background()
	s, err := cli.Device(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.BatteryPercentage)
}

```

## Features
- Add Test For client.go
- Obtain Sesame locking history from client
