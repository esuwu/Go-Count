package open

import (
	"io"
	"net/http"
	"os"
	"strings"
)

type ClientT struct {
	httpClient *http.Client
}

func (c *ClientT) open(name string) (io.ReadCloser, error) {
	if strings.HasPrefix(name, "http://") || strings.HasPrefix(name, "https://") {
		resp, err := c.httpClient.Get(name)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
	return os.Open(name)
}

func Open(name string) (io.ReadCloser, error) {
	c := &ClientT{httpClient: http.DefaultClient}
	return c.open(name)
}
