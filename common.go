package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/daolicloud/daolictl/cli"
	"github.com/daolicloud/daolictl/opts"
)

const (
	defaultCaFile	= "ca.pem"
	defaultKeyFile 	= "key.pem"
	defaultCertFile = "cert.pem"
	tlsverify	= "tlsverify"
)

var (
	commonFlags = &cli.CommonFlags{FlagSet: new(flag.FlagSet)}

        daoliCertPath  = os.Getenv("DOCKER_CERT_PATH")
        daoliTLSVerify = os.Getenv("DOCKER_TLS_VERIFY") != ""
)

func init() {
	commonFlags.PostParse = postParseCommon
	cmd := commonFlags.FlagSet

	cmd.BoolVar(&commonFlags.Debug, "D,-debug", false, "Enable debug mode")
        cmd.StringVar(&commonFlags.LogLevel, "l,-log-level", "info", "Set the logging level")
        cmd.BoolVar(&commonFlags.TLS, "-tls", false, "Use TLS; implied by --tlsverify")
        cmd.BoolVar(&commonFlags.TLSVerify, "-tlsverify", daoliTLSVerify, "Use TLS and verify the remote")

        // TODO use flag flag.String("i,-identity", "", "Path to libtrust key file")
        var tlsOptions cli.Options
        commonFlags.TLSOptions = &tlsOptions
        cmd.StringVar(&tlsOptions.CAFile, "-tlscacert", filepath.Join(daoliCertPath, defaultCaFile), "Trust certs signed only by this CA")
        cmd.StringVar(&tlsOptions.CertFile, "-tlscert", filepath.Join(daoliCertPath, defaultCertFile), "Path to TLS certificate file")
        cmd.StringVar(&tlsOptions.KeyFile, "-tlskey", filepath.Join(daoliCertPath, defaultKeyFile), "Path to TLS key file")

	validateHost := func(val string) (string, error) {return val, nil }
        cmd.Var(opts.NewNamedListOptsRef("hosts", &commonFlags.Hosts, validateHost), "H,-host", "Daemon socket(s) to connect to")
}

func postParseCommon() {
        // Regardless of whether the user sets it to true or false, if they
        // specify --tlsverify at all then we need to turn on tls
        // TLSVerify can be true even if not set due to DOCKER_TLS_VERIFY env var, so we need to check that here as well
        if commonFlags.TLSVerify {
                commonFlags.TLS = true
        }

        if !commonFlags.TLS {
                commonFlags.TLSOptions = nil
        } else {
                tlsOptions := commonFlags.TLSOptions
                tlsOptions.InsecureSkipVerify = !commonFlags.TLSVerify
                tlsOptions.CertFile = ""
                tlsOptions.KeyFile = ""
        }

}
