package telegram

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client{
	return Client{
		host: host,
		basePath: newBasePath(token),
		client: http.Client{},
	}
}

func newBasePath(token string) string{
	return "bot"+token
}


func (c *Client) Updates(offset int, limit int)([]Update, error){
	 q := url.Values{}
	 q.Add("offset", strconv.Itoa(offset))
	 q.Add("limit", strconv.Itoa(limit))

	 // do request <- GetUpdates
}

func (c *Client) doRequst(method string, query url.Values) ([]byte, error){
	u:=url.URL{
		Scheme: "https",
		Host: c.host,
		Path: c.basePath + method,
	}

}

func (c *Client) SendMessage (){

}