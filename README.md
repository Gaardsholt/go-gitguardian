# go-gitguardian - GitGuardian API Client

Go API client library for the [GitGuardian API](https://api.gitguardian.com/)

## How to use this library

Each category of endpoints has it own package.

So if you want to use the `scanning` endpoint, you import the `scan` package like this:

```go
package main

import (
	"fmt"
	"log"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/scan"
)

func main() {

	c, err := scan.NewClient(client.WithApiKey("<YOUR_API_KEY>"))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.ContentScan(scan.ContentScanPayload{
		Filename: "my_file.py",
		Document: `import urllib.request
    url = 'http://jen_barber:correcthorsebatterystaple@cake.gitguardian.com/isreal.json'
    response = urllib.request.urlopen(url)
    consume(response.read())`,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp.Result)
}
```

You can set the API key by either using the `client.WithApiKey` function or by setting the environment varible `GITGUARDIAN_API_KEY`.

If you are running a self hosted version of GitGuardian you can also set the server address by either using the `client.WithServer` function or by setting the environment varible `GITGUARDIAN_SERVER`.


## Stuff missing from this library

  * Better logging
  * Do a checkup on the error handling
  * Add tests
  * Properly more...
