package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/daolicloud/daolictl/client/transport"
)

type Client struct {
	// proto holds the client protocol i.e. unix.
	proto string
	// addr holds the client address.
	addr string
	// basePath holds the path to prepend to the requests.
	basePath string
	// transport is the interface to sends request with, it implements transport.Client.
	transport transport.Client
	// version of the server to talk to.
	version string
	// custom http headers configured by usrs.
	customHTTPHeaders map[string]string
}

// NewClient initialize a new API client for the given host and API version.
func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error) {
	proto, addr, basePath, err := ParseHost(host)
	if err != nil {
		return nil, err
	}

	transport, err := transport.NewTransportWithHTTP(proto, addr, client)
	if err != nil {
		return nil, err
	}

	return &Client{
		proto:             proto,
		addr:              addr,
		basePath:          basePath,
		transport:         transport,
		version:           version,
		customHTTPHeaders: httpHeaders,
	}, nil
}

// getAPIPath returns the versioned request path to call the api.
// It appends the query parameters to the path if they are not empty.
func (cli *Client) getAPIPath(p string, query url.Values) string {
	var apiPath string
	if cli.version != "" {
		v := strings.TrimPrefix(cli.version, "v")
		apiPath = fmt.Sprintf("%s/v%s%s", cli.basePath, v, p)
	} else {
		apiPath = fmt.Sprintf("%s%s", cli.basePath, p)
	}
	if len(query) > 0 {
		apiPath += "?" + query.Encode()
	}
	return apiPath
}

// ParseHost verifies that that the given host strings is valid.
func ParseHost(host string) (string, string, string, error) {
	protoAddrParts := strings.SplitN(host, "://", 2)
	if len(protoAddrParts) == 1 {
		return "", "", "", fmt.Errorf("unable to parse daolinet host `%s`", host)
	}

	var basePath string
	proto, addr := protoAddrParts[0], protoAddrParts[1]
	if proto == "tcp" {
		parsed, err := url.Parse("tcp://" + addr)
		if err != nil {
			return "", "", "", err
		}
		addr = parsed.Host
		basePath = parsed.Path
	}
	return proto, addr, basePath, nil

}
