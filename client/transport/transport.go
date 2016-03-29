package transport

import (
	"fmt"
	"net/http"

	"github.com/docker/go-connections/sockets"
)

type cliTransport struct {
	*http.Client
	*tlsInfo
	transport *http.Transport
}

func NewTransportWithHTTP(proto, addr string, client *http.Client) (Client, error) {
	var transport *http.Transport

	if client != nil {
		tr, ok := client.Transport.(*http.Transport)
		if !ok {
			return nil, fmt.Errorf("unable to verify TLS configuration, invalid transport %v", client.Transport)
		}
		transport = tr
	} else {
		transport = defaultTransport(proto, addr)
		client = &http.Client{
			Transport: transport,
		}
	}

	return &cliTransport{
		Client:    client,
		tlsInfo:   &tlsInfo{transport.TLSClientConfig},
		transport: transport,
	}, nil
}

// CancelRequest stops a request execution.
func (a *cliTransport) CancelRequest(req *http.Request) {
	a.transport.CancelRequest(req)
}

// defaultTransport creates a new http.Transport with Docker's
// default transport configuration.
func defaultTransport(proto, addr string) *http.Transport {
	tr := new(http.Transport)
	sockets.ConfigureTransport(tr, proto, addr)
	return tr
}

var _ Client = &cliTransport{}
