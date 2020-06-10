package client

import (
        "github.com/looqbox/statuspage-go/pkg/root"
)

func Connect(pageId string, apiKey string) root.BaseRequest {
        request := root.BaseRequest{
                Url: "https://api.statuspage.io/v1/pages/" + pageId,
                Headers: []root.Header{
                        root.Header{
                                Name:  "Authorization",
                                Value: "OAuth " + apiKey,
                        },
                },
        }

        return request
}