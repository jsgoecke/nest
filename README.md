#nest
[![wercker status](https://app.wercker.com/status/caae45a6a0f785001032fac0c6775f3c/m "wercker status")](https://app.wercker.com/project/bykey/caae45a6a0f785001032fac0c6775f3c)

A Go library for the [Nest](http://developer.nest.com) API for Nest devices. This is early work and only supports querying the devices object as well as the REST Streaming API for devices.

## Version

0.0.1

## Installation

	go get github.com/jsgoecke/nest

## Documentation

[http://godoc.org/github.com/jsgoecke/nest](http://godoc.org/github.com/jsgoecke/nest)

## Usage

```go
package main

import (
	"../."
	"fmt"
)

const (
	ClientID          = "<client-id>"
	State             = "STATE"
	ClientSecret      = "<client-secret>"
	AuthorizationCode = "<authorization-code> - https://developer.nest.com/documentation/how-to-auth"
)

func main() {
	client := nest.New(ClientID, State, ClientSecret, AuthorizationCode)
	client.Authorize()
	client.DevicesStream(func(event *nest.Devices) {
		fmt.Println(event)
	})
}
```

## Testing
	
	cd nest
	go test

## License

MIT, see LICENSE.txt

## Author

Jason Goecke [@jsgoecke](http://twitter.com/jsgoecke)

## Todo

Per the write permissions here:

[https://developer.nest.com/documentation/api](https://developer.nest.com/documentation/api)

Provide additional functions to update the following settings:

### Structures

	* away
	* ETA support
