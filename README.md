## Sesame
This package provides a client to the Sesame API.

This package API Client to lock the keys of Sesame API using signatures encrypted with the AES-CMAC method.

For more information about the API, please see the official documentation.

https://doc.candyhouse.co/ja/SesameAPI

## Example
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

## Features
- Add Test For client.go
- Obtain Sesame information from the client
- Obtain Sesame locking history from client
