package statuspageio_test

import (
	"fmt"

	"github.com/looqbox/statuspage-go"
)

func ExampleConnect() {
	pageId := "tLwrSv54NJCJWHYUzLl3n0dMngEc9Cr8"
	apiKey := "5OCMVqruAKqXsZpT5bbrs3ISgsgMu01Z"

	request := statuspageio.Connect(pageId, apiKey)
	fmt.Println(request)
	// Output:
	// {https://api.statuspage.io/v1/pages/tLwrSv54NJCJWHYUzLl3n0dMngEc9Cr8 [{Authorization OAuth 5OCMVqruAKqXsZpT5bbrs3ISgsgMu01Z}]}
}