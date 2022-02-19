package main

import (
	"fmt"
	"log"

	"github.com/Gaardsholt/go-gitguardian/scan"
)

func main() {

	c, err := scan.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.MultipleContentScan([]scan.ContentScanPayload{
		{
			Document: `
	import urllib.request
	url = 'http://jen_barber:correcthorsebatterystaple@cake.gitguardian.com/isreal.json'
	response = urllib.request.urlopen(url)
	consume(response.read())`,
			Filename: "asd",
		}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp.Result)
}
