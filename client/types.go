package client

type (
	Firewall struct {
		Name        string
		Container   string
		DatapathID  string
		GatewayIP   string
		GatewayPort int
		ServicePort int
	}
)
