package myanimelist

import "github.com/nstratos/go-myanimelist/mal"

type Client struct {
	Client *mal.Client
}

func New() *Client {
	return &Client{}
}
