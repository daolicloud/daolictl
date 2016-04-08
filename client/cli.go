package client

import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/daolicloud/daolictl/cli"
	"github.com/docker/go-connections/sockets"
)

const (
	DefaultVersion = "1.23"
	DefaultHost    = "tcp://127.0.0.1:3380"
)

// DaoliCli represents the daolictl command line client.
// instances of the client can be returned from NewDaoliCli.
type DaoliCli struct {
	// initializing closure
	init func() error

	// in holds the input stream and closer (io.ReadCloser) for the client.
	in io.ReadCloser
	// out holds the output stream (io.Writer) for the client.
	out io.Writer
	// err holds the error stream (io.Writer) for the client.
	err io.Writer
	// client is the http client that performs all API operations
	client APIClient
}

// Initialize calls the init function that will setup the configuration for the client
// such as the TLS, tcp and other parameters used to run the client.
func (cli *DaoliCli) Initialize() error {
	if cli.init == nil {
		return nil
	}
	return cli.init()
}

func NewDaoliCli(in io.ReadCloser, out, err io.Writer, clientFlags *cli.ClientFlags) *DaoliCli {
	cli := &DaoliCli{
		in:  in,
		out: out,
		err: err,
	}

	cli.init = func() error {
		clientFlags.PostParse()
		host, err := getServerHost(clientFlags.Common.Hosts, clientFlags.Common.TLSOptions)
		if err != nil {
			return err
		}

		customHeaders := map[string]string{}
		customHeaders["User-Agent"] = "Daolictl-Client/ (" + runtime.GOOS + ")"

		verStr := DefaultVersion
		if tmpStr := os.Getenv("DAOLI_API_VERSION"); tmpStr != "" {
			verStr = tmpStr
		}
		httpClient, err := newHTTPClient(host, clientFlags.Common.TLSOptions)
		if err != nil {
			return err
		}

		client, err := NewClient(host, verStr, httpClient, customHeaders)
		if err != nil {
			return err
		}
		cli.client = client

		return nil
	}

	return cli
}

func getServerHost(hosts []string, tlsOptions *cli.Options) (host string, err error) {
	switch len(hosts) {
	case 0:
		host = os.Getenv("DAOLI_HOST")
	case 1:
		host = hosts[0]
	default:
		return "", errors.New("Please specify only one -H")
	}

	host = strings.TrimSpace(host)
	if host == "" {
		host = DefaultHost
	}

	return
}

func newHTTPClient(host string, tlsOptions *cli.Options) (*http.Client, error) {
	if tlsOptions == nil {
		// let the api client configure the default transport.
		return nil, nil
	}

	var config = &tls.Config{}
	tr := &http.Transport{
		TLSClientConfig: config,
	}
	proto, addr, _, err := ParseHost(host)
	if err != nil {
		return nil, err
	}

	sockets.ConfigureTransport(tr, proto, addr)

	return &http.Client{
		Transport: tr,
	}, nil
}
